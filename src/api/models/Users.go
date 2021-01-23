package models

import (
	"../security"
	"errors"
	"github.com/badoux/checkmail"
	"gorm.io/gorm"
	"html"
	"strings"
	"time"
)

type User struct {
	ID uint64 			`gorm:"primaryKey;autoIncrement" json:"id"`
	Nickname string		`gorm:"size:20;not null;unique" json:"nickname"`
	Email string		`gorm:"size:50;not null;unique" json:"email"`
	Password string		`gorm:"size:60;not null" json:"password,omitempty"` //the field is omitted from the object if its value is empty,
	CreatedAt time.Time	`gorm:"DEFAULT CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time	`gorm:"DEFAULT CURRENT_TIMESTAMP" json:"updated_at"`
	Posts []Post `gorm:"foreignKey:AuthorID;constraint:OnUpdate:CASCADE" json:"posts,omitempty"` //the field is omitted from the object if its value is empty,
}

func (u *User) BeforeSave(tx *gorm.DB) error  {
	hashedPassword, err := security.Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}


func (u *User) Prepare(tx *gorm.DB) {
	u.ID = 0
	u.Nickname = html.EscapeString(strings.TrimSpace(u.Nickname))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

func (u *User) Validate(tx *gorm.DB) error {
	if u.Nickname == "" {
		return errors.New("required nickname")
	}
	if u.Password == "" {
		return errors.New("required password")
	}
	if u.Email == "" {
		return errors.New("required email")
	}
	if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("invalid email")
	}
	return nil
}
