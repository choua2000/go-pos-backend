package models

import (
	"gorm.io/gorm"
)

type ProductImage struct {
	gorm.Model
	ProductID uint
	URL       string
	PublicID  string
	Product   Product `gorm:"foreignKey:ProductID"`
}
