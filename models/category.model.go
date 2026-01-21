package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name     string    `gorm:"size:100;not null"`
	Products []Product `gorm:"foreignKey:CategoryID"`
}
