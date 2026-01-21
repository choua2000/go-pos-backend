package utils

import (
	"go-backend/config" // Adjust this to your project name
	"go-backend/models"
)

func CheckEmailExists(email string) bool {
	var user models.User
	// .Select("id") is faster than selecting the whole row
	result := config.DB.Where("email = ?", email).First(&user)

	// If RowsAffected is > 0, the email exists
	return result.RowsAffected > 0
}
