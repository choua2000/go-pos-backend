package services

import (
	"go-backend/config"
	"go-backend/models"
)

func GetAllSaleItems() ([]models.SaleItem, error) {
	var saleItems []models.SaleItem
	result := config.DB.
		Preload("Product").
		Find(&saleItems)
	return saleItems, result.Error
}

func CreateSaleItem(saleItem *models.SaleItem) error {
	result := config.DB.Create(saleItem)
	return result.Error
}

func GetSaleItemByID(id string) (*models.SaleItem, error) {
	var saleItem models.SaleItem
	result := config.DB.
		Preload("Product").
		First(&saleItem, id)
	return &saleItem, result.Error
}

func GetSaleItemsBySaleID(saleID uint) ([]models.SaleItem, error) {
	var saleItems []models.SaleItem
	result := config.DB.
		Preload("Product").
		Where("sale_id = ?", saleID).
		Find(&saleItems)
	return saleItems, result.Error
}

func UpdateSaleItem(id string, updates map[string]interface{}) (*models.SaleItem, error) {
	var saleItem models.SaleItem
	result := config.DB.Model(&models.SaleItem{}).
		Where("id = ?", id).
		Updates(updates)
	if result.Error != nil {
		return nil, result.Error
	}
	config.DB.
		Preload("Product").
		First(&saleItem, id)
	return &saleItem, nil
}

func DeleteSaleItem(id string) error {
	result := config.DB.Delete(&models.SaleItem{}, id)
	return result.Error
}

func ValidateSaleID(id uint) error {
	return config.DB.First(&models.Sale{}, id).Error
}

func ValidateProductID(id uint) error {
	return config.DB.First(&models.Product{}, id).Error
}
