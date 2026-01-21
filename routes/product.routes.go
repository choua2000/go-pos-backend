package routes

import (
	"go-backend/controllers"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(r *gin.Engine) {
	product := r.Group("/api/products")
	{
		product.GET("/getAll", controllers.GetProducts)
		product.POST("/create", controllers.CreateProduct)
		product.GET("/get/:id", controllers.GetProductByID)
		product.PUT("/update/:id", controllers.UpdateProductByID)
		product.DELETE("/delete/:id", controllers.DeleteProductByID)
	}
}
