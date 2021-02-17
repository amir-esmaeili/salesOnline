package models
// Database Configs

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func SetUpDataBase() {
	// Load environment variables
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	dbPath := os.Getenv("DB")

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// Migrate
	err = db.AutoMigrate(&User{})
	if err != nil {
		panic(err)
	}

	// Using database as a global object
	DB = db
}
