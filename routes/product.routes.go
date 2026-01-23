package routes

import (
	"go-backend/controllers"
	"go-backend/middlewares"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(r *gin.Engine) {
	product := r.Group("/api/products")
	{
		product.GET("/getAll", controllers.GetProducts)
		product.POST("/create", middlewares.AuthMiddleware(), middlewares.AdminOnly(), controllers.CreateProduct)
		product.GET("/get/:id", controllers.GetProductByID)
		product.PUT("/update/:id", middlewares.AuthMiddleware(), middlewares.AdminOnly(), controllers.UpdateProductByID)
		product.DELETE("/delete/:id", middlewares.AuthMiddleware(), middlewares.AdminOnly(), controllers.DeleteProductByID)
		product.POST("/images/:id/images", controllers.UploadProductImages)
	}
}
