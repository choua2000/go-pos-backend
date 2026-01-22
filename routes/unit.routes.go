package routes

import (
	"go-backend/controllers"
	"go-backend/middlewares"

	"github.com/gin-gonic/gin"
)

func UnitRoutes(r *gin.Engine) {
	unit := r.Group("/api/units")
	{
		unit.GET("/getAll", controllers.GetUnits)
		unit.POST("/create", middlewares.AuthMiddleware(), middlewares.AdminOnly(), controllers.CreateUnit)
		unit.GET("/get/:id", controllers.GetUnitByID)
		unit.PUT("/update/:id", middlewares.AuthMiddleware(), middlewares.AdminOnly(), controllers.UpdateUnitByID)
		unit.DELETE("/delete/:id", middlewares.AuthMiddleware(), middlewares.AdminOnly(), controllers.DeleteUnitByID)
	}
}
