package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Subscribe(c *gin.Context) {
	// Lógica para suscribir a un curso
	c.JSON(http.StatusOK, gin.H{"message": "subscribe successful"})
}
