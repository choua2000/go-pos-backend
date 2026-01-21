package services

import (
	"errors"
	"go-backend/config"
	"go-backend/models"
	"go-backend/utils"
)

var (
	ErrEmailExists = errors.New("email already exists")
)

func RegisterUser(user *models.User) error {

	// 1. Check email exists
	if utils.CheckEmailExists(user.Email) {
		return ErrEmailExists
	}

	// 2. Hash password
	hash, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hash

	// 3. Create user
	if err := config.DB.Create(user).Error; err != nil {
		return err
	}

	return nil
}
