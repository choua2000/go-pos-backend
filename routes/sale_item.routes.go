package routes

import (
	"go-backend/controllers"

	"github.com/gin-gonic/gin"
)

func SaleItemRoutes(router *gin.Engine) {
	saleItemGroup := router.Group("/api/sale-items")
	{
		saleItemGroup.GET("/getAll", controllers.GetSaleItems)
		saleItemGroup.POST("/create", controllers.CreateSaleItem)
		saleItemGroup.GET("/get/:id", controllers.GetSaleItemByID)
		saleItemGroup.GET("/sale/:sale_id", controllers.GetSaleItemsBySaleID)
		saleItemGroup.PUT("/update/:id", controllers.UpdateSaleItemByID)
		saleItemGroup.DELETE("/delete/:id", controllers.DeleteSaleItemByID)
	}
}
