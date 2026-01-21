package utils

import (
	"go-backend/models"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

var JwtKey []byte

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: No .env file found, using system environment variables")
	}

	secret := os.Getenv("SECRET_KEY")
	if secret == "" {
		log.Fatal("SECRET_KEY is not set")
	}

	JwtKey = []byte(secret)
}

func GenerateToken(userID uint, email string, name string, role string) (string, error) {
	claims := models.JwtClaims{
		ID:    userID,
		Email: email,
		Name:  name,
		Role:  role, // admin / user
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(JwtKey))
	return tokenString, err

}
