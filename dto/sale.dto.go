package dto

type CreateSaleRequest struct {
	InvoiceNumber string        `json:"invoice_number" binding:"required"`
	UserID        uint          `json:"user_id" binding:"required"`
	TotalAmount   float64       `json:"total_amount" binding:"required"`
	PaymentMethod string        `json:"payment_method" binding:"required,oneof=cash transfer"`
	SaleItems     []SaleItemDTO `json:"sale_items" binding:"required"`
}

type UpdateSaleRequest struct {
	InvoiceNumber string        `json:"invoice_number"`
	UserID        uint          `json:"user_id"`
	TotalAmount   float64       `json:"total_amount"`
	PaymentMethod string        `json:"payment_method" binding:"oneof=cash transfer"`
	SaleItems     []SaleItemDTO `json:"sale_items"`
}

type SaleItemDTO struct {
	ProductID uint    `json:"product_id" binding:"required"`
	Quantity  int     `json:"quantity" binding:"required,gt=0"`
	UnitPrice float64 `json:"unit_price" binding:"required,gt=0"`
}
