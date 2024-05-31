package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	// Lógica de autenticación
	c.JSON(http.StatusOK, gin.H{"message": "login successful"})
}
