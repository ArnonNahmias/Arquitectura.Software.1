package controllers

import (
	"backend/domain"
	"backend/services"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

/*func Search(c *gin.Context) {
	// Lógica de búsqueda de cursos
	c.JSON(http.StatusOK, gin.H{"message": "search successful"})
}
*/

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
