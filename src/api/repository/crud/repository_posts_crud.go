package crud

import (
	"../../models"
	"../../utils/channels"
	"errors"
	_ "errors"
	"gorm.io/gorm"
	"time"
	_ "time"
)

type repositoryPostsCRUD struct {
	db *gorm.DB
}

func NewRepositoryPostsCRUD(db *gorm.DB) *repositoryPostsCRUD {
	return &repositoryPostsCRUD{db}
}

func (r *repositoryPostsCRUD) Save(post models.Post) (models.Post, error) {
	var err error
	done := make(chan bool)

	go func(ch chan<- bool) {
		err = r.db.Debug().Model(&models.Post{}).Create(&post).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return post, nil
	}
	return models.Post{}, err
}
func (r *repositoryPostsCRUD) FindAll() ([]models.Post, error) {
	var err error
	done := make(chan bool)
	var posts []models.Post
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Debug().Model(&models.Post{}).Limit(100).Find(&posts).Error
		if err != nil {
			ch <- false
			return
		}
		if len(posts) > 0 {
			for i := range posts {
				err = r.db.Debug().Model(&models.User{}).Where("id = ?", posts[i].AuthorID).Take(&posts[i].Author).Error
				if err != nil {
					ch <- false
					return
				}
			}
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return posts, nil
	}
	return nil, err
}
func (r *repositoryPostsCRUD) FindById(pid uint64) (models.Post, error) {
	var err error
	done := make(chan bool)
	var post models.Post
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Debug().Model(&models.Post{}).Where("id = ?", pid).Take(&post).Error
		if err != nil {
			ch <- false
			return
		}
		if post.ID != 0 {

			err = r.db.Debug().Model(&models.User{}).Where("id = ?", post.AuthorID).Take(&post.Author).Error
			if err != nil {
				ch <- false
				return
			}
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return post, nil
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.Post{}, errors.New("post not found")
	}
	return models.Post{}, err
}

func (r *repositoryPostsCRUD) Update(pid uint64, post models.Post) (int64, error) {
	var rs  *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		rs = r.db.Debug().Model(&models.Post{}).Where("id = ?", pid).UpdateColumns(
			map[string]interface{}{
				"title": post.Title,
				"content": post.Content,
				"updated_at": time.Now(),
			},
		)
		ch <- true
	}(done)
	if channels.OK(done) {
		if rs.Error != nil {
			if errors.Is(rs.Error, gorm.ErrRecordNotFound) {
				return 0, errors.New("post not found")
			}
			return 0, rs.Error
		}
		return rs.RowsAffected, nil
	}
	return 0, rs.Error
}

func (r *repositoryPostsCRUD) Delete(uid uint64) (int64, error) {
	var rs  *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		rs = r.db.Debug().Model(&models.Post{}).Where("id = ?", uid).Delete(&models.Post{})
		ch <- true
	}(done)
	if channels.OK(done) {
		if  rs.Error != nil {
			return 0, rs.Error
		}
		return rs.RowsAffected, nil
	}
	return 0, rs.Error
}