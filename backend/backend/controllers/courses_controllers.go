package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Get(c *gin.Context) {
	// Lógica para obtener un curso por ID
	c.JSON(http.StatusOK, gin.H{"message": "get course successful"})
}
