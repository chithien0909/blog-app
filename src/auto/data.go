package auto

import (
	"../api/models"
)

var users = []models.User{
	{
		Nickname: "Jhon Doe",
		Email: "jhondoe@email.com",
		Password: "1234566789",
	},
}

var posts = []models.Post{
	{
		Title:     "Title post 01",
		Content:   "Post 01 content",
	},
}