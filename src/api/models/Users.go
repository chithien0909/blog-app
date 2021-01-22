package models

import (
	"../security"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID int64 			`gorm:"primaryKey;autoIncrement" json:"id"`
	Nickname string		`gorm:"size:20;not null;unique" json:"nickname"`
	Email string		`gorm:"size:50;not null;unique" json:"email"`
	Password string		`gorm:"size:60;not null" json:"password,omitempty"`
	CreatedAt time.Time	`gorm:"DEFAULT CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time	`gorm:"DEFAULT CURRENT_TIMESTAMP" json:"updated_at"`

}

func (u *User) BeforeSave(tx *gorm.DB) error  {
	hashedPassword, err := security.Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}