package initializers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetupDB() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := os.Getenv("DB_URL")

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connect to DB")
	}
}

func SetupDBTesting() {
	err := os.Chdir("..")

	if err != nil {
		log.Fatal("Failed to change directory: ", err)
	}

	err = godotenv.Load(".env")
	if err != nil {
		log.Fatal("Failed to load .env file: ", err)
	}

	dsn := os.Getenv("DB_URL")

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connect to DB")
	}
}
