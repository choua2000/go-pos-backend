package controllers

import (
	"go-backend/config"
	"go-backend/models"
	"go-backend/services"
	"go-backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {

	var body models.User

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid data",
		})
		return
	}

	err := services.RegisterUser(&body)
	if err != nil {

		if err == services.ErrEmailExists {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Register failed",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"code":    200,
		"status":  "success",
		"message": "Register successfully",
		"data":    body,
	})
}

func Login(c *gin.Context) {
	var body models.User
	var user models.User

	c.ShouldBindJSON(&body)
	config.DB.Where("email = ?", body.Email).First(&user)

	if user.ID == 0 || !utils.CheckPassword(user.Password, body.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, _ := utils.GenerateToken(user.ID, user.Email, user.Name, user.Role)
	c.JSON(http.StatusOK, gin.H{"code": "200", "status": "success", "message": "Login successfully", "token": token})
}
