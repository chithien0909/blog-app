package crud

import (
	"../../models"
	"../../utils/channels"
	"errors"
	"gorm.io/gorm"
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