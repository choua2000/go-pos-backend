package models

import "gorm.io/gorm"

type SaleItem struct {
	gorm.Model
	SaleID    uint    `json:"saleId"`
	ProductID uint    `json:"productId"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price" gorm:"type:decimal(10,2)"`
	Subtotal  float64 `json:"subtotal" gorm:"type:decimal(10,2)"`
}
