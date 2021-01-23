package auto

import (
	"../api/models"
)

var users = []models.User{
	{
		Nickname: "Jhon Doe",
		Email: "jhondoe@email.com",
		Password: "123456",
	},
	{
		Nickname: "Chi Thien",
		Email: "chithien@gmail.com",
		Password: "123456",
	},
}

var posts = []models.Post{
	{
		Title:     "Title post 01",
		Content:   "Post 01 content",
	},
	{
		Title:     "Title post 02",
		Content:   "Post 02 content",
	},
}