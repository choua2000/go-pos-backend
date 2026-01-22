package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"size:100;not null"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Role     string `gorm:"type:varchar(20);default:'admin'"` // admin or cashier
	Sales    []Sale `gorm:"foreignKey:UserID"`
}
