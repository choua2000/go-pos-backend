package routes

import (
	"go-backend/controllers"

	"github.com/gin-gonic/gin"
)

func SaleRoutes(router *gin.Engine) {
	saleGroup := router.Group("/api/sales")
	{
		saleGroup.GET("/gets", controllers.GetSales)
		saleGroup.POST("/create", controllers.CreateSale)
		saleGroup.GET("/get/:id", controllers.GetSaleByID)
		// saleGroup.PUT("/update/:id", controllers.UpdateSaleByID)
		// saleGroup.DELETE("/delete/:id", controllers.DeleteSaleByID)
	}
}
