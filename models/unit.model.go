package models

import (
	"gorm.io/gorm"
)

type Unit struct {
	gorm.Model
	Name string `gorm:"size:20;not null"` // pcs, kg, box
}
