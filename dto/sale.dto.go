package dto

type SaleItemInput struct {
	ProductID uint `json:"productId" binding:"required"`
	Quantity  int  `json:"quantity" binding:"required,min=1"`
}

type CreateSaleRequest struct {
	UserID        uint            `json:"userId" binding:"required"`
	PaymentMethod string          `json:"paymentMethod" binding:"required,oneof=cash transfer"`
	Items         []SaleItemInput `json:"items" binding:"required,min=1"`
}
