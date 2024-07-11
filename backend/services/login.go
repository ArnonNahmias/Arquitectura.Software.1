// services/login.go
package services

import (
	"errors"
	"fmt"
	"time"

	"backend/clients"
	"backend/dao"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var jwtKey = []byte("my_secret_key")

type Claims struct {
	Username string `json:"username"`
	UserID   int    `json:"userId"`
	jwt.StandardClaims
}

func GenerateJWT(username string, userID int) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Username: username,
		UserID:   userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func Login(username, password string) (string, int, string, error) {
	fmt.Println("Username:", username) // Depuración
	fmt.Println("Password:", password) // Depuración

	var user dao.Usuario
	if err := clients.DB.Where("nombre_usuario = ?", username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Println("User not found") // Depuración
			return "", 0, "", errors.New("invalid credentials")
		}
		return "", 0, "", err
	}

	fmt.Println("User found:", user.NombreUsuario) // Depuración
	fmt.Println("Hashed Password in DB:", user.Contrasena) // Depuración

	err := bcrypt.CompareHashAndPassword([]byte(user.Contrasena), []byte(password))
	if err != nil {
		fmt.Println("Invalid password") // Depuración
		return "", 0, "", errors.New("invalid credentials")
	}

	token, err := GenerateJWT(username, user.IdUsuario)
	if err != nil {
		return "", 0, "", err
	}

	return token, user.IdUsuario, user.Tipo, nil
}

func ValidateJWT(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, err
		}
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
