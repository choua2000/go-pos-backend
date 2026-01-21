package controllers

import (
	"go-backend/dto"
	"go-backend/models"
	"go-backend/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// MEANS GET ALL SALES
func GetSales(c *gin.Context) {
	sales, err := services.GetAllSales()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch sales"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": "200", "status": "success", "message": "Get sales successfully", "data": sales})
}

// MEAN : CREATE NEW SALE
func CreateSale(c *gin.Context) {
	var req dto.CreateSaleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	if err := services.ValidateUserID(req.UserID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	sale := models.Sale{
		InvoiceNumber: req.InvoiceNumber,
		UserID:        req.UserID,
		TotalAmount:   req.TotalAmount,
		PaymentMethod: req.PaymentMethod,
	}

	if err := services.CreateSale(&sale); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create sale"})
		return
	}

	result, _ := services.GetSaleByID(strconv.FormatUint(uint64(sale.ID), 10))
	c.JSON(http.StatusCreated, gin.H{"code": "200", "status": "success", "message": "Create sale successfully", "data": result})
}

// MEAN : GET SALE BY ID
func GetSaleByID(c *gin.Context) {
	id := c.Param("id")
	sale, err := services.GetSaleByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Sale not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": "200", "status": "success", "message": "Get sale successfully", "data": sale})
}

// MEAN : UPDATE SALE BY ID
func UpdateSaleByID(c *gin.Context) {
	id := c.Param("id")
	_, err := services.GetSaleByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "Sale not found"})
		return
	}

	var req dto.UpdateSaleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}

	if req.UserID != 0 {
		if err := services.ValidateUserID(req.UserID); err != nil {
			c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
			return
		}
	}

	updates := map[string]interface{}{}
	if req.InvoiceNumber != "" {
		updates["invoice_number"] = req.InvoiceNumber
	}
	if req.UserID != 0 {
		updates["user_id"] = req.UserID
	}
	if req.TotalAmount != 0 {
		updates["total_amount"] = req.TotalAmount
	}
	if req.PaymentMethod != "" {
		updates["payment_method"] = req.PaymentMethod
	}

	result, err := services.UpdateSale(id, updates)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to update sale"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": "200", "status": "success", "message": "Update sale successfully", "data": result})
}

// MEAN : DELETE SALE BY ID
func DeleteSaleByID(c *gin.Context) {
	id := c.Param("id")
	_, err := services.GetSaleByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Sale not found"})
		return
	}

	if err := services.DeleteSale(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete sale"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": "200", "status": "success", "message": "Delete sale successfully"})
}
