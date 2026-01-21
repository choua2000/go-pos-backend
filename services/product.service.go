package services

import (
	"go-backend/config"
	"go-backend/models"
)

func GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	result := config.DB.
		Joins("LEFT JOIN categories ON categories.id = products.category_id").
		Joins("LEFT JOIN units ON units.id = products.unit_id").
		Preload("Category").
		Preload("Unit").
		Where("categories.id IS NOT NULL AND units.id IS NOT NULL").
		Find(&products)
	return products, result.Error
}

func CreateProduct(product *models.Product) error {
	result := config.DB.Create(product)
	return result.Error
}

func GetProductByID(id string) (*models.Product, error) {
	var product models.Product
	result := config.DB.
		Preload("Category").
		Preload("Unit").
		First(&product, id)
	return &product, result.Error
}

func UpdateProduct(id string, updates map[string]interface{}) (*models.Product, error) {
	var product models.Product
	result := config.DB.Model(&models.Product{}).
		Where("id = ?", id).
		Updates(updates)
	if result.Error != nil {
		return nil, result.Error
	}
	config.DB.
		Preload("Category").
		Preload("Unit").
		First(&product, id)
	return &product, nil
}

func DeleteProduct(id string) error {
	result := config.DB.Delete(&models.Product{}, id)
	return result.Error
}

func ValidateCategoryID(id uint) error {
	return config.DB.First(&models.Category{}, id).Error
}

func ValidateUnitID(id uint) error {
	return config.DB.First(&models.Unit{}, id).Error
}
