package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

var (
	PORT = 0
	DB_DRIVER string
	DB_URL string
	SECRET_KEY []byte
)

func Load() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Load env failed: %v\n", err)
	}
	PORT, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		PORT = 9000
	}
	DB_DRIVER = os.Getenv("DB_DRIVER")
	DB_URL = fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"))
	SECRET_KEY = []byte(os.Getenv("API_SECRET"))
}