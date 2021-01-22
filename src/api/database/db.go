package database

import (
	"../../config"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"

)

func Connect() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(config.DB_URL), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil

}
