package database

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func InitDatabase() {
	err := godotenv.Load("backend/.env")
	if err != nil {
		log.Print("Error loading .env file")
	}
	DB, err = gorm.Open(postgres.Open(os.Getenv("POSTGRES_CONN")), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
	log.Print("Database connection established")
}
