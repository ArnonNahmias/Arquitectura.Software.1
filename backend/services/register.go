// services/register.go
package services

import (
	"backend/clients"
	"backend/dao"
	"golang.org/x/crypto/bcrypt"
)

func Register(username, password, userType string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := dao.Usuario{
		NombreUsuario: username,
		Contrasena:    string(hashedPassword),
		Tipo:          userType,
	}

	if err := clients.DB.Create(&user).Error; err != nil {
		return err
	}

	return nil
}
