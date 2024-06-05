package controllers

import (
	"backend/domain"
	"backend/services"
	"net/http"

	"fmt"

	"github.com/gin-gonic/gin"
)

/*func Subscribe(c *gin.Context) {
	// Lógica para suscribir a un curso
	c.JSON(http.StatusOK, gin.H{"message": "subscribe successful"})
}
*/

func Subscribe(c *gin.Context) {
	var request domain.SubscribeRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, domain.Result{
			Message: fmt.Sprintf("Invalid request: %s", err.Error()),
		})
		return
	}

	if err := services.Subscribe(request.UserID, request.CourseID); err != nil {
		c.JSON(http.StatusConflict, domain.Result{
			Message: fmt.Sprintf("Error in subscribe; %s", err.Error()),
		})
		return
	}

	c.JSON(http.StatusCreated, domain.Result{
		Message: fmt.Sprintf("User %d successfully subscribed to course %d", request.UserID, request.CourseID),
	})
}
