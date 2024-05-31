package courses

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/yourproject/domain"
	"github.com/yourusername/yourproject/services"
)

// Search busca cursos según una consulta.
func Search(c *gin.Context) {
	query := strings.TrimSpace(c.Query("query"))
	results, err := services.Search(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.Result{
			Message: fmt.Sprintf("Error in search: %s", err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, domain.SearchResponse{
		Results: results,
	})
}

// Get obtiene un curso por su ID.
func Get(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.Result{
			Message: fmt.Sprintf("Invalid ID: %s", err.Error()),
		})
		return
	}

	course, err := services.GetCourse(id)
	if err != nil {
		c.JSON(http.StatusNotFound, domain.Result{
			Message: fmt.Sprintf("Error in get: %s", err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, course)
}

// Subscribe permite suscribirse a un curso.
func Subscribe(c *gin.Context) {
	var request domain.SubscribeRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, domain.Result{
			Message: "Invalid request: " + err.Error(),
		})
		return
	}

	// Lógica para suscribir al usuario al curso
	// ...

	c.JSON(http.StatusOK, domain.Result{
		Message: "User subscribed successfully!",
	})
}
