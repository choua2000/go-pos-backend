package main

import (
	"go-backend/config"
	"go-backend/routes"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load env
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect DB
	config.ConnectDB()

	r := gin.Default()

	// ===== Secure CORS from ENV =====
	origins := strings.Split(os.Getenv("CLIENT_ORIGINS"), ",")

	r.Use(cors.New(cors.Config{
		AllowOrigins: origins,
		AllowMethods: []string{
			"GET", "POST", "PUT", "DELETE", "OPTIONS",
		},
		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Authorization",
			"Accept",
			"X-Requested-With",
		},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	// =================================

	// Routes
	routes.AuthRoutes(r)
	routes.UnitRoutes(r)
	routes.CategoryRoutes(r)
	routes.ProductRoutes(r)
	routes.SaleRoutes(r)
	routes.SaleItemRoutes(r)
	routes.ProductImageRoutes(r)

	r.GET("/api", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Go API is running 🚀",
		})
	})

	PORT := os.Getenv("PORT")

	r.Run(":" + PORT)
}
