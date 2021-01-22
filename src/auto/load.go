package auto

import (
	"../api/database"
	_ "../api/utils/console"
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

	// Drop table if exists
	//err = db.Debug().Migrator().DropTable(&models.User{})
	//if err != nil {
	//	log.Fatalf("Drop table if existed failed: %v\n", err)
	//}
	// Create table, missing foreign keys, constraints, columns and indexes
	// It WON’T delete unused columns to protect your data.
	//err = db.Debug().AutoMigrate(&models.User{})
	//if err != nil {
	//	log.Fatalf("Drop table if existed failed: %v\n", err)
	//}
	//for _, user := range users {
	//	err = db.Debug().Model(&models.User{}).Create(&user).Error
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	console.Pretty(user)
	//}
}