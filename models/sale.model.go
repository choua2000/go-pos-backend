package models

import "gorm.io/gorm"

type Sale struct {
	gorm.Model
	InvoiceNumber string     `json:"invoiceNumber" gorm:"unique;not null"`
	UserID        uint       `json:"userId"`
	User          User       `json:"user" gorm:"foreignKey:UserID"`
	TotalAmount   float64    `json:"totalAmount" gorm:"type:decimal(10,2)"`
	PaymentMethod string     `json:"paymentMethod" gorm:"size:50"` // cash, transfer
	SaleItems     []SaleItem `json:"saleItems" gorm:"foreignKey:SaleID"`
}
