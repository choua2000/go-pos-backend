package routes

import (
	"go-backend/controllers"
	"go-backend/middlewares"

	"github.com/gin-gonic/gin"
)

func CategoryRoutes(r *gin.Engine) {
	category := r.Group("/api/categories")
	{
		category.GET("/getAll", controllers.GetCategories)
		category.POST("/create", middlewares.AuthMiddleware(), middlewares.AdminOnly(), controllers.CreateCategory)
		category.GET("/get/:id", controllers.GetCategoryByID)
		category.PUT("/update/:id", middlewares.AuthMiddleware(), middlewares.AdminOnly(), controllers.UpdateCategoryByID)
		category.DELETE("/delete/:id", controllers.DeleteCategoryByID)
	}
}
