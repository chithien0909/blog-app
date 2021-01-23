package repository

import "../models"
type UserRepository interface {
	Save(models.User) (models.User, error)
	FindAll() ([]models.User, error)
	FindById(uint64) (models.User, error)
	Update(uint64, models.User) (int64, error)
	Delete(uint64) (int64, error)
}