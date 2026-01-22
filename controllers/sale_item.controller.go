package controllers

import (
	"go-backend/dto"
	// "go-backend/models"
	"go-backend/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// MEANS GET ALL SALE ITEMS
func GetSaleItems(c *gin.Context) {
	saleItems, err := services.GetAllSaleItems()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch sale items"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": "200", "status": "success", "message": "Get sale items successfully", "data": saleItems})
}

// MEAN : CREATE NEW SALE ITEM
func CreateSaleItem(c *gin.Context) {

	var req dto.CreateSaleItemRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	saleItem, err := services.CreateSaleItem(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"code":    "200",
		"status":  "success",
		"message": "Sale item created successfully",
		"data":    saleItem,
	})
}

// MEAN : GET SALE ITEM BY ID
func GetSaleItemByID(c *gin.Context) {
	id := c.Param("id")
	saleItem, err := services.GetSaleItemByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Sale item not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": "200", "status": "success", "message": "Get sale item successfully", "data": saleItem})
}

// MEAN : GET SALE ITEMS BY SALE ID
func GetSaleItemsBySaleID(c *gin.Context) {
	saleID := c.Param("sale_id")
	saleIDUint := uint(0)
	if _, err := strconv.ParseUint(saleID, 10, 32); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid sale ID"})
		return
	}
	saleIDUint = uint(0)
	val, _ := strconv.ParseUint(saleID, 10, 32)
	saleIDUint = uint(val)

	saleItems, err := services.GetSaleItemsBySaleID(saleIDUint)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch sale items"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": "200", "status": "success", "message": "Get sale items successfully", "data": saleItems})
}

// MEAN : UPDATE SALE ITEM BY ID
func UpdateSaleItemByID(c *gin.Context) {
	id := c.Param("id")
	_, err := services.GetSaleItemByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "Sale item not found"})
		return
	}

	var req dto.UpdateSaleItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}

	if req.SaleID != 0 {
		if err := services.ValidateSaleID(req.SaleID); err != nil {
			c.JSON(http.StatusNotFound, gin.H{"message": "Sale not found"})
			return
		}
	}

	if req.ProductID != 0 {
		if err := services.ValidateProductID(req.ProductID); err != nil {
			c.JSON(http.StatusNotFound, gin.H{"message": "Product not found"})
			return
		}
	}

	updates := map[string]interface{}{}
	if req.SaleID != 0 {
		updates["sale_id"] = req.SaleID
	}
	if req.ProductID != 0 {
		updates["product_id"] = req.ProductID
	}
	if req.Quantity != 0 {
		updates["quantity"] = req.Quantity
	}
	if req.Price != 0 {
		updates["price"] = req.Price
	}
	if req.Subtotal != 0 {
		updates["subtotal"] = req.Subtotal
	}

	result, err := services.UpdateSaleItem(id, updates)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to update sale item"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": "200", "status": "success", "message": "Update sale item successfully", "data": result})
}

// MEAN : DELETE SALE ITEM BY ID
func DeleteSaleItemByID(c *gin.Context) {
	id := c.Param("id")
	_, err := services.GetSaleItemByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Sale item not found"})
		return
	}

	if err := services.DeleteSaleItem(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete sale item"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": "200", "status": "success", "message": "Delete sale item successfully"})
}
