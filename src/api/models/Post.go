package models

import "time"

type Post struct {
	ID uint64 			`gorm:"primaryKey;autoIncrement" json:"id"`
	Title string 		`gorm:"size:30;not null;unique" json:"title"`
	Content string	`gorm:"size:255;not null;unique" json:"content"`
	Author User			`json:"author"` //Skip this field for json encoding
	AuthorID uint64 `gorm:"not null" json:"author_id"`
	CreatedAt time.Time	`gorm:"default currentTimestamp" json:"created_at"`
	UpdatedAt time.Time	`gorm:"default currentTimestamp" json:"updated_at"`
}
