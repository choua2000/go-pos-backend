package main

import (
	"go-backend/config"
	"go-backend/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	config.ConnectDB()
	r := gin.Default()
	routes.AuthRoutes(r)
	routes.UnitRoutes(r)
	routes.CategoryRoutes(r)
	routes.ProductRoutes(r)
	routes.SaleRoutes(r)
	routes.SaleItemRoutes(r)
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	PORT := os.Getenv("PORT")
	r.GET("/api", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Go API is running 🚀",
		})
	})

	r.Run(":" + PORT)
}
