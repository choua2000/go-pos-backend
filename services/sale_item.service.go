package services

import (
	"errors"
	"go-backend/config"
	"go-backend/dto"
	"go-backend/models"
)

func GetAllSaleItems() ([]models.SaleItem, error) {
	var saleItems []models.SaleItem
	result := config.DB.
		Preload("Product").
		Find(&saleItems)
	return saleItems, result.Error
}

func CreateSaleItem(req dto.CreateSaleItemRequest) (*models.SaleItem, error) {
	if err := ValidateSaleID(req.SaleID); err != nil {
		return nil, errors.New("sale not found with id: " + string(rune(req.SaleID)))
	}

	if err := ValidateProductID(req.ProductID); err != nil {
		return nil, errors.New("product not found with id: " + string(rune(req.ProductID)))
	}

	saleItem := models.SaleItem{
		SaleID:    req.SaleID,
		ProductID: req.ProductID,
		Quantity:  req.Quantity,
		Price:     req.Price,
		Subtotal:  req.Subtotal,
	}

	result := config.DB.Create(&saleItem)
	if result.Error != nil {
		return nil, result.Error
	}

	config.DB.Preload("Product").First(&saleItem, saleItem.ID)
	return &saleItem, nil
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
