package controllers

import (
	"fmt"
	"go-backend/config"
	"go-backend/helper"
	"go-backend/models"
	"go-backend/services"
	"go-backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UploadProductImages(c *gin.Context) {
	productID := c.Param("id")

	// Validate product ID is not empty
	if productID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product ID is required"})
		return
	}

	// Validate product ID format
	pid, err := helper.StringToUint(productID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID format"})
		return
	}

	// Verify product exists
	var product models.Product
	if err := config.DB.First(&product, pid).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse form"})
		return
	}

	files := form.File["images"]

	if len(files) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Images required"})
		return
	}

	var images []models.ProductImage
	var uploadErrors []string

	for _, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			uploadErrors = append(uploadErrors, fmt.Sprintf("Failed to open file: %s", fileHeader.Filename))
			continue
		}
		defer file.Close()

		buffer, err := utils.ResizeAndCompress(file)
		if err != nil {
			uploadErrors = append(uploadErrors, fmt.Sprintf("Failed to compress: %s", err.Error()))
			continue
		}

		url, publicID, err := services.UploadToCloudinary(buffer)
		if err != nil {
			uploadErrors = append(uploadErrors, fmt.Sprintf("Failed to upload: %s", err.Error()))
			continue
		}

		image := models.ProductImage{
			ProductID: pid,
			URL:       url,
			PublicID:  publicID,
		}
		images = append(images, image)
	}

	if len(images) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "No images uploaded successfully",
			"details": uploadErrors,
		})
		return
	}

	if err := config.DB.Create(&images).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save images to database", "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"code":    "200",
		"status":  "success",
		"message": "Upload images success",
		"data":    images,
	})
}

func DeleteProductWithImages(c *gin.Context) {
	productID := c.Param("id")

	// Validate product ID is not empty
	if productID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product ID is required"})
		return
	}

	tx := config.DB.Begin()

	var images []models.ProductImage
	tx.Where("product_id = ?", productID).Find(&images)

	for _, img := range images {
		_ = services.DeleteFromCloudinary(img.PublicID)
	}

	if err := tx.Where("product_id = ?", productID).
		Delete(&models.ProductImage{}).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete images"})
		return
	}

	if err := tx.Delete(&models.Product{}, productID).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
		return
	}

	if err := tx.Commit().Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to commit transaction"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    "200",
		"status":  "success",
		"message": "Delete product and images successfully",
	})
}
