package auth

import (
	"../database"
	"../models"
	"../security"
	"../utils/channels"
	"database/sql"
	"fmt"
	"gorm.io/gorm"
)


func SignIn(email string, password string) (string, error) {

	user := models.User{}
	var err error
	var db *gorm.DB
	var mysql *sql.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		db, err = database.Connect()
		if err != nil {
			ch<- false
			return
		}
		mysql, err = db.DB()
		if err != nil {
			ch<- false
			return
		}
		defer mysql.Close()

		err = db.Debug().Model(&models.User{}).Where("email = ?", email).Take(&user).Error
		if err != nil {
			ch<- false
			return
		}
		fmt.Println(user.Password)
		fmt.Println(password)
		err = security.VerifyPassword(user.Password, password)
		if err != nil {
			ch<- false
			return
		}
		ch<- true
	}(done)
	if channels.OK(done) {

		return CreateToken(user.ID)
	}
	return "", err
}