// models/jwt_claims.go
package models

import "github.com/golang-jwt/jwt/v5"

type JwtClaims struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
	Role  string `json:"role"` // 👈 this is missing in RegisteredClaims
	jwt.RegisteredClaims
}
