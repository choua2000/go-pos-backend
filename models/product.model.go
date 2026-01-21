package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name       string   `gorm:"size:255;not null"`
	Price      float64  `gorm:"type:decimal(10,2);not null"`
	Stock      int      `gorm:"not null"`
	CategoryID uint     `gorm:"not null"`
	UnitID     uint     `gorm:"not null"`
	Category   Category `gorm:"foreignKey:CategoryID"`
	Unit       Unit     `gorm:"foreignKey:UnitID"`
}
