package migrations

import (
	"go-backend/models"

	"gorm.io/gorm"
)

func CreateProductImagesTable(db *gorm.DB) error {
	return db.AutoMigrate(&models.ProductImage{})
}
