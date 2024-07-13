package controllers

import (
    "backend/clients"
    "backend/dao"
    "backend/domain"
    "backend/services"
    "log"
    "net/http"
    "github.com/gin-gonic/gin"
)

func CreateSubscription(c *gin.Context) {
    var request domain.SubscribeRequest
    if err := c.ShouldBindJSON(&request); err != nil {
        log.Printf("Error binding JSON: %v", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
        return
    }

    // Logging the received request for debugging
    log.Printf("Received subscription request: %+v\n", request)

    // Validate IDs are not zero
    if request.UserID == 0 || request.CourseID == 0 {
        log.Printf("Invalid user ID or course ID: UserID=%d, CourseID=%d", request.UserID, request.CourseID)
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID or course ID"})
        return
    }

    subscription := dao.Subscription{
        IdUsuario: int(request.UserID),
        IdCurso:   int(request.CourseID),
    }

    log.Printf("Creating subscription with User ID: %d and Course ID: %d", subscription.IdUsuario, subscription.IdCurso)
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

func GetSubscriptions(c *gin.Context) {
	var subscriptions []dao.Subscription
	if err := clients.DB.Find(&subscriptions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching subscriptions"})
		return
	}
	c.JSON(http.StatusOK, subscriptions)
}

func DeleteSubscription(c *gin.Context) {
	var subscription dao.Subscription
	id := c.Param("id")
	if err := clients.DB.First(&subscription, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Subscription not found"})
		return
	}
	if err := clients.DB.Delete(&subscription).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting subscription"})
		return
	}
	c.Status(http.StatusNoContent)
}
