package config

import (
	"go-backend/models"

	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	dsn := os.Getenv("DB_URL")
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	DB = database
	database.AutoMigrate(&models.User{})
	database.AutoMigrate(&models.Product{})
	database.AutoMigrate(&models.Unit{})
	database.AutoMigrate(&models.Sale{})
	database.AutoMigrate(&models.SaleItem{})
	database.AutoMigrate(&models.Category{})
}
