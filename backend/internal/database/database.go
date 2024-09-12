package database

import (
	"backend/internal/models"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func InitDatabase() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	// Открытие соединения с базой данных через GORM
	DB, err = gorm.Open(postgres.Open(os.Getenv("POSTGRES_CONN")), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	err = DB.AutoMigrate(&models.Tender{}, &models.Bid{})
	if err != nil {
		log.Println("Failed to migrate database:", err)
	}
}
