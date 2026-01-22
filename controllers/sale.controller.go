package controllers

import (
	"go-backend/dto"
	"go-backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateSale(c *gin.Context) {

	var req dto.CreateSaleRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sale, err := services.CreateSale(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"code":    "200",
		"status":  "true",
		"message": "Sale created successfully",
		"data":    sale,
	})
}

// MEANS GET ALL SALES
func GetSales(c *gin.Context) {
	sales, err := services.GetSales()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch sales"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": "200", "status": "true", "message": "Get sales successfully", "data": sales})
}

// MEAN : GET SALE BY ID
func GetSaleByID(c *gin.Context) {
	id := c.Param("id")
	sale, err := services.GetSaleByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Sale not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": "200", "status": "true", "message": "Get sale by id successfully", "data": sale})
}
