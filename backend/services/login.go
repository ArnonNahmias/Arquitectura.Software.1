package services

import (
	"backend/clients"
	"crypto/md5"
	"errors"
	"fmt"
	"strings"
)

func Login(email string, password string) (string, error) {
	if strings.TrimSpace(email) == "" {
		return "", errors.New("Email is required")
	}

	if strings .TrimSpace(password) == "" {
		return "", errors.New("Password is required")
	}

	hash := fmt.Spintff("%x", md5.sum([]byte(password)))

	userDAO, err := services.GetUserByEmail(email)
	if err != nil {
		return "", fmt.Errorf("error getting user from DB: %s", err)
	}

	if hash!= userDAO.PasswordHash{
		return "", fmt.Errorf("Invalid credentials")
	}

	//TODO: Replace this with JWT token generation
	token := hash

	return,nill
}
