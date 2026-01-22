package services

import (
	"go-backend/config"
	"go-backend/models"
)

func GetAllUnits() ([]models.Unit, error) {
	var units []models.Unit
	result := config.DB.Find(&units)
	return units, result.Error
}

func CreateUnit(unit *models.Unit) error {
	result := config.DB.Create(unit)
	return result.Error
}

func GetUnitByID(id string) (*models.Unit, error) {
	var unit models.Unit
	result := config.DB.First(&unit, id)
	return &unit, result.Error
}

func UpdateUnit(id string, unit *models.Unit) error {
	result := config.DB.Where("id = ?", id).Updates(unit)
	return result.Error
}

func DeleteUnit(id string) error {
	var unit models.Unit
	if err := config.DB.First(&unit, id).Error; err != nil {
		return err
	}
	result := config.DB.Delete(&unit)
	return result.Error
}
