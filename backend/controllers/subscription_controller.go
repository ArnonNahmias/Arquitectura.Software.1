package controllers

import (
	"log"
	"backend/domain"
	"backend/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateSubscription creates a new subscription.
func CreateSubscription(c *gin.Context) {
	var request domain.SubscribeRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	// Logging the received request
	log.Printf("Received subscription request: %+v\n", request)

	// Validate IDs are not zero
	if request.UserID == 0 || request.CourseID == 0 {
		log.Println("Invalid user ID or course ID")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID or course ID"})
		return
	}

	subscription := domain.Subscription{
		IdUsuario: request.UserID,
		IdCurso:   request.CourseID,
	}

	log.Println("Creating subscription...")
	newSubscription, err := services.CreateSubscription(subscription)
	if err != nil {
		log.Printf("Error creating subscription: %v", err)
		if err.Error() == "user not found" || err.Error() == "course not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating subscription", "details": err.Error()})
		}
		return
	}
	log.Println("Subscription created successfully")
	c.JSON(http.StatusCreated, newSubscription)
}

// GetSubscriptions retrieves all subscriptions.
func GetSubscriptions(c *gin.Context) {
	subscriptions, err := services.GetSubscriptions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching subscriptions"})
		return
	}
	c.JSON(http.StatusOK, subscriptions)
}

// DeleteSubscription deletes a subscription.
func DeleteSubscription(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid subscription ID"})
		return
	}

	if err := services.DeleteSubscription(id); err != nil {
		if err.Error() == "subscription not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Subscription not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting subscription"})
		}
		return
	}
	c.Status(http.StatusNoContent)
}
