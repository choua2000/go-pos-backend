package dto

type CreateSaleItemRequest struct {
	SaleID    uint    `json:"saleId" binding:"required"`
	ProductID uint    `json:"productId" binding:"required"`
	Quantity  int     `json:"quantity" binding:"required,min=1"`
	Price     float64 `json:"price" binding:"omitempty,gt=0"`
	Subtotal  float64 `json:"subtotal" binding:"omitempty,gt=0"`
}

type UpdateSaleItemRequest struct {
	SaleID    uint    `json:"sale_id"`
	ProductID uint    `json:"product_id"`
	Quantity  int     `json:"quantity" binding:"omitempty,gt=0"`
	Price     float64 `json:"price" binding:"omitempty,gt=0"`
	Subtotal  float64 `json:"subtotal" binding:"omitempty,gt=0"`
}
