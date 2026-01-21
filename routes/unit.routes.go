package routes

import (
	"go-backend/controllers"

	"github.com/gin-gonic/gin"
)

func UnitRoutes(r *gin.Engine) {
	unit := r.Group("/api/units")
	{
		unit.GET("/getAll", controllers.GetUnits)
		unit.POST("/create", controllers.CreateUnit)
		unit.GET("/get/:id", controllers.GetUnitByID)
		unit.PUT("/update/:id", controllers.UpdateUnitByID)
		unit.DELETE("/delete/:id", controllers.DeleteUnitByID)
	}
}
