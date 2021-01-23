package auto

import (
	"../api/database"
	"../api/models"
	"../api/utils/console"
	"log"
)

func Load()  {
	db, err := database.Connect()
	if err != nil {
		log.Fatalf("Error while connecting database: %v\n", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Error while calling sqlDB: %v\n", err)
	}
	defer sqlDB.Close()

	//Drop table if exists
	err = db.Debug().Migrator().DropTable(&models.Post{},&models.User{})
	if err != nil {
		log.Fatalf("Drop table if existed failed: %v\n", err)
	}
	// Create table, missing foreign keys, constraints, columns and indexes
	// It WON’T delete unused columns to protect your data.
	err = db.Debug().AutoMigrate(&models.Post{}, &models.User{})
	if err != nil {
		log.Fatalf("Drop table if existed failed: %v\n", err)
	}
	//err := db.Debug().Debug().Key
	for i := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatal(err)
		}
		console.Pretty(users[i])
		posts[i].AuthorID = users[i].ID
		posts[i].Author = users[i]
		err = db.Debug().Model(&models.Post{}).Create(&posts[i]).Error
		if err != nil {
			log.Fatal(err)
		}
		console.Pretty(posts[i])
	}
}
