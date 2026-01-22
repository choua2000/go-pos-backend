package services

import (
	"errors"
	"fmt"
	"go-backend/config"
	"go-backend/dto"
	"go-backend/models"
	"time"

	"gorm.io/gorm"
)

func CreateSale(req dto.CreateSaleRequest) (*models.Sale, error) {

	var saleItems []models.SaleItem
	var totalAmount float64

	err := config.DB.Transaction(func(tx *gorm.DB) error {

		for _, item := range req.Items {

			var product models.Product

			// 1️⃣ Get product
			if err := tx.First(&product, item.ProductID).Error; err != nil {
				return errors.New("product not found")
			}

			// 2️⃣ Check stock
			if product.Stock < item.Quantity {
				return errors.New("not enough stock")
			}

			// 3️⃣ Calculate subtotal
			subtotal := product.Price * float64(item.Quantity)

			// 4️⃣ Reduce stock
			product.Stock -= item.Quantity
			if err := tx.Save(&product).Error; err != nil {
				return err
			}

			// 5️⃣ Add sale item
			saleItems = append(saleItems, models.SaleItem{
				ProductID: product.ID,
				Quantity:  item.Quantity,
				Price:     product.Price,
				Subtotal:  subtotal,
			})

			// 6️⃣ Add to total
			totalAmount += subtotal
		}

		// 7️⃣ Create sale
		sale := models.Sale{
			InvoiceNumber: generateInvoice(),
			UserID:        req.UserID,
			PaymentMethod: req.PaymentMethod,
			TotalAmount:   totalAmount,
			SaleItems:     saleItems,
		}

		if err := tx.Create(&sale).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	var sale models.Sale
	config.DB.Preload("SaleItems").Last(&sale)

	return &sale, nil
}

func generateInvoice() string {
	return fmt.Sprintf("INV-%d", time.Now().Unix())
}

// GetSales retrieves all sales with their associated sale items.
func GetSales() ([]models.Sale, error) {
	var sales []models.Sale
	result := config.DB.Preload("SaleItems").Find(&sales)
	return sales, result.Error
}
