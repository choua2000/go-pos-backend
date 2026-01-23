package middlewares

import (
	"go-backend/models"
	"go-backend/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Missing token"})
			c.Abort()
			return
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

		token, err := utils.ValidateToken(tokenString)
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
			c.Abort()
			return
		}

		claims := token.Claims.(*models.JwtClaims)

		// Save user info to context
		c.Set("user_id", claims.ID)
		c.Set("role", claims.Role)

		c.Next()
	}
}

func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {

		role, exists := c.Get("role")
		if !exists || role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{
				"message": "Admin access only",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

// MEAN : Cashier ONLY

func CashierOnly() gin.HandlerFunc {
	return func(c *gin.Context) {

		role, exists := c.Get("role")
		if !exists || role != "cashier" {
			c.JSON(http.StatusForbidden, gin.H{
				"message": "Cashier access only",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
