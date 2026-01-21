package dto

type CreateSaleItemRequest struct {
	SaleID    uint    `json:"sale_id" binding:"required"`
	ProductID uint    `json:"product_id" binding:"required"`
	Quantity  int     `json:"quantity" binding:"required,gt=0"`
	Price     float64 `json:"price" binding:"required,gt=0"`
	Subtotal  float64 `json:"subtotal" binding:"required,gt=0"`
}

type UpdateSaleItemRequest struct {
	SaleID    uint    `json:"sale_id"`
	ProductID uint    `json:"product_id"`
	Quantity  int     `json:"quantity" binding:"gt=0"`
	Price     float64 `json:"price" binding:"gt=0"`
	Subtotal  float64 `json:"subtotal" binding:"gt=0"`
}
