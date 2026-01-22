package controllers

import (
	"go-backend/models"
	"go-backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// MEANS GET ALL CATEGORIES
func GetCategories(c *gin.Context) {
	categories, err := services.GetAllCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch categories"})
		return
	}
	c.JSON(http.StatusOK, categories)
}

// MEAN : CREATE NEW CATEGORY
func CreateCategory(c *gin.Context) {
	var body models.Category
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}
	if err := services.CreateCategory(&body); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create category"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"code": "200", "status": "success", "message": "Create category successfully", "data": body})
}

// MEAN : GET CATEGORY BY ID
func GetCategoryByID(c *gin.Context) {
	id := c.Param("id")
	category, err := services.GetCategoryByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": "200", "status": "success", "message": "Get category successfully", "data": category})
}

// MEAN : UPDATE CATEGORY BY ID
func UpdateCategoryByID(c *gin.Context) {
	id := c.Param("id")
	category, err := services.GetCategoryByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}
	if err := services.UpdateCategory(id, category); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update category"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": "200", "status": "success", "message": "Update category successfully", "data": category})
}

// MEAN : DELETE CATEGORY BY ID
func DeleteCategoryByID(c *gin.Context) {
	id := c.Param("id")
	_, err := services.GetCategoryByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}
	if err := services.DeleteCategory(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete category"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": "200", "status": "success", "message": "Delete category successfully"})
}
