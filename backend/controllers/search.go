package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Search(c *gin.Context) {
	// Lógica de búsqueda de cursos
	c.JSON(http.StatusOK, gin.H{"message": "search successful"})
}
