package models

import (
	"errors"
	"html"
	"strings"
	"time"
)

type Post struct {
	ID uint64 			`gorm:"primaryKey;autoIncrement" json:"id"`
	Title string 		`gorm:"size:30;not null;unique" json:"title"`
	Content string	`gorm:"size:255;not null;unique" json:"content"`
	Author User			`gorm:"_" json:"author"` //Skip this field for json encoding
	AuthorID uint64 `gorm:"not null" json:"author_id"`
	CreatedAt time.Time	`gorm:"default currentTimestamp" json:"created_at"`
	UpdatedAt time.Time	`gorm:"default currentTimestamp" json:"updated_at"`
}

//func (u *Post) BeforeSave(tx *gorm.DB) error  {
//	err := tx.Model(&User{}).Where("id = ?", u.AuthorID).Take(&u.Author).Error
//	if err != nil {
//		return errors.New("author not found")
//	}
//	return nil
//}

func (p *Post) Prepare()  {
	p.ID = 0
	p.Title = html.EscapeString(strings.TrimSpace(p.Title))
	p.Content = html.EscapeString(strings.TrimSpace(p.Content))
	p.Author = User{}
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()

}
func (p *Post) Validate(action string) error {

	switch strings.ToLower(action) {
		case "update":
			if p.Title == "" {
				return errors.New("required title")
			}
			if p.Content == "" {
				return errors.New("required content")
			}
			return nil
		default:
			if p.Title == "" {
				return errors.New("required title")
			}
			if p.Content == "" {
				return errors.New("required content")
			}
			if p.AuthorID < 1 {
				return errors.New("required AuthorID")
			}
			return nil
	}
	return nil
}