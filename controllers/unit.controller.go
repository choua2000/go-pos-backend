package controllers

import (
	"go-backend/models"
	"go-backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// MEANS GET ALL UNITS
func GetUnits(c *gin.Context) {
	units, err := services.GetAllUnits()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch units"})
		return
	}
	c.JSON(http.StatusOK, units)
}

// MEAN : CREATE NEW UNIT
func CreateUnit(c *gin.Context) {
	var body models.Unit
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}
	if err := services.CreateUnit(&body); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create unit"})
		return
	}
	c.JSON(http.StatusCreated, body)
}

// MEAN : GET UNIT BY ID
func GetUnitByID(c *gin.Context) {
	id := c.Param("id")
	unit, err := services.GetUnitByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unit not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": "200", "status": "success", "message": "Get unit successfully", "data": unit})
}

// MEAN : UPDATE UNIT BY ID
func UpdateUnitByID(c *gin.Context) {
	id := c.Param("id")
	unit, err := services.GetUnitByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unit not found"})
		return
	}

	if err := c.ShouldBindJSON(&unit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}
	if err := services.UpdateUnit(id, unit); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update unit"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": "200", "status": "success", "message": "Update unit successfully", "data": unit})
}

// MEAN : DELETE UNIT BY ID
func DeleteUnitByID(c *gin.Context) {
	id := c.Param("id")
	unit, err := services.GetUnitByID(id)
	if err != nil || unit.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unit not found"})
		return
	}
	if err := services.DeleteUnit(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete unit"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": "200", "status": "success", "message": "Delete unit successfully"})
}
