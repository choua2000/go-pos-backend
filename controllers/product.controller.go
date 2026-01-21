package controllers

import (
	"go-backend/dto"
	"go-backend/models"
	"go-backend/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// MEANS GET ALL PRODUCTS
func GetProducts(c *gin.Context) {
	products, err := services.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": "200", "status": "success", "message": "Get products successfully", "data": products})
}

// MEAN : CREATE NEW PRODUCT
func CreateProduct(c *gin.Context) {
	var req dto.CreateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	if err := services.ValidateCategoryID(req.CategoryID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	if err := services.ValidateUnitID(req.UnitID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid unit ID"})
		return
	}

	product := models.Product{
		Name:       req.Name,
		Price:      req.Price,
		Stock:      req.Stock,
		CategoryID: req.CategoryID,
		UnitID:     req.UnitID,
	}

	if err := services.CreateProduct(&product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}

	result, _ := services.GetProductByID(strconv.FormatUint(uint64(product.ID), 10))
	c.JSON(http.StatusCreated, gin.H{"code": "200", "status": "success", "message": "Create product successfully", "data": result})
}

// MEAN : GET PRODUCT BY ID
func GetProductByID(c *gin.Context) {
	id := c.Param("id")
	product, err := services.GetProductByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": "200", "status": "success", "message": "Get product successfully", "data": product})
}

// MEAN : UPDATE PRODUCT BY ID
func UpdateProductByID(c *gin.Context) {
	id := c.Param("id")
	_, err := services.GetProductByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "Product not found"})
		return
	}

	var req dto.UpdateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}

	if err := services.ValidateCategoryID(req.CategoryID); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Category not found"})
		return
	}

	if err := services.ValidateUnitID(req.UnitID); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Unit not found"})
		return
	}

	updates := map[string]interface{}{
		"name":        req.Name,
		"price":       req.Price,
		"stock":       req.Stock,
		"category_id": req.CategoryID,
		"unit_id":     req.UnitID,
	}

	result, err := services.UpdateProduct(id, updates)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to update product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": "200", "status": "success", "message": "Update product successfully", "data": result})
}

// MEAN : DELETE PRODUCT BY ID
func DeleteProductByID(c *gin.Context) {
	id := c.Param("id")
	_, err := services.GetProductByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Product not found"})
		return
	}

	if err := services.DeleteProduct(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": "200", "status": "success", "message": "Delete product successfully"})
}
