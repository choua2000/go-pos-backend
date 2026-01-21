package utils

import (
	"go-backend/models"

	"github.com/golang-jwt/jwt/v5"
)

func ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(
		tokenString,
		&models.JwtClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return JwtKey, nil
		},
	)
}
