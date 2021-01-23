package crud

import (
	"../../models"
	"../../utils/channels"
	"errors"
	"gorm.io/gorm"
	"time"
)

type repositoryUsersCRUD struct {
	db *gorm.DB
}

func NewRepositoryUsersCRUD(db *gorm.DB) *repositoryUsersCRUD {
	return &repositoryUsersCRUD{db}
}

func (r *repositoryUsersCRUD) Save(user models.User) (models.User, error) {
	var err error
	done := make(chan bool)

	go func(ch chan<- bool) {
		err = r.db.Debug().Model(&models.User{}).Create(&user).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return user, nil
	}
	return models.User{}, err
}
func (r *repositoryUsersCRUD) FindAll() ([]models.User, error) {
	var err error
	done := make(chan bool)
	var users []models.User
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Debug().Model(&models.User{}).Limit(100).Find(&users).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return users, nil
	}
	return nil, err
}
func (r *repositoryUsersCRUD) FindById(uid uint64) (models.User, error) {
	var err error
	done := make(chan bool)
	var user models.User
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Debug().Model(&models.User{}).Where("id = ?", uid).Take(&user).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return user, nil
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.User{}, errors.New("user not found")
	}
	return models.User{}, err
}

func (r *repositoryUsersCRUD) Update(uid uint64, user models.User) (int64, error) {
	var rs  *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		rs = r.db.Debug().Model(&models.User{}).Where("id = ?", uid).UpdateColumns(
			map[string]interface{}{
				"nickname": user.Nickname,
				"email": user.Email,
				"updated_at": time.Now(),
			},
		)
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

func (r *repositoryUsersCRUD) Delete(uid uint64) (int64, error) {
	var rs  *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		rs = r.db.Debug().Model(&models.User{}).Where("id = ?", uid).Delete(&models.User{})
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