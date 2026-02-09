package config

import (
	"log"
	"os"

	"github.com/joho/godotenv" // <- tambahkan ini
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	// Load file .env
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, relying on OS env variables")
	}

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL is empty")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB connection failed:", err)
	}

	DB = db
	log.Println("DB connected")
}
