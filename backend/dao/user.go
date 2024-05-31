package dao

import "time"

// User representa un usuario en el sistema.
type user struct {
	ID           int64     // ID del usuario
	EmailHash    string    // Hash del correo electrónico del usuario
	PasswordHash string    // Hash de la contraseña del usuario
	Type         string    // Tipo de usuario (admin o normal)
	CreationDate time.Time // Fecha de creación del usuario
	LastUpdated  time.Time // Fecha de última actualización del usuario
}
