package services

import (
	"go-backend/config"
	"go-backend/models"
)

func GetAllSales() ([]models.Sale, error) {
	var sales []models.Sale
	result := config.DB.
		Preload("User").
		Preload("SaleItems").
		Find(&sales)
	return sales, result.Error
}

func CreateSale(sale *models.Sale) error {
	result := config.DB.Create(sale)
	return result.Error
}

func GetSaleByID(id string) (*models.Sale, error) {
	var sale models.Sale
	result := config.DB.
		Preload("User").
		Preload("SaleItems").
		First(&sale, id)
	return &sale, result.Error
}

func UpdateSale(id string, updates map[string]interface{}) (*models.Sale, error) {
	var sale models.Sale
	result := config.DB.Model(&models.Sale{}).
		Where("id = ?", id).
		Updates(updates)
	if result.Error != nil {
		return nil, result.Error
	}
	config.DB.
		Preload("User").
		Preload("SaleItems").
		First(&sale, id)
	return &sale, nil
}

func DeleteSale(id string) error {
	result := config.DB.Delete(&models.Sale{}, id)
	return result.Error
}

func ValidateUserID(id uint) error {
	return config.DB.First(&models.User{}, id).Error
}
