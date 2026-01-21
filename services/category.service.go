package services

import (
	"go-backend/config"
	"go-backend/models"
)

func GetAllCategories() ([]models.Category, error) {
	var categories []models.Category
	result := config.DB.Find(&categories)
	return categories, result.Error
}

func CreateCategory(category *models.Category) error {
	result := config.DB.Create(category)
	return result.Error
}

func GetCategoryByID(id string) (*models.Category, error) {
	var category models.Category
	result := config.DB.First(&category, id)
	return &category, result.Error
}

func UpdateCategory(id string, category *models.Category) error {
	result := config.DB.Where("id = ?", id).Updates(category)
	return result.Error
}

func DeleteCategory(id string) error {
	result := config.DB.Delete(&models.Category{}, id)
	return result.Error
}
