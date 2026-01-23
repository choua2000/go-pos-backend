package routes

import (
	"go-backend/controllers"

	"github.com/gin-gonic/gin"
)

func ProductImageRoutes(r *gin.Engine) {
	productImage := r.Group("/api/product-images")
	{
		productImage.POST("/upload", controllers.UploadProductImages)
		productImage.DELETE("/delete/:id", controllers.DeleteProductWithImages)
	}
}
