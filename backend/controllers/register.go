package controllers


import (
    "net/http"
    "backend/services"
    "github.com/gin-gonic/gin"
)

func RegisterC(c *gin.Context) {
    var userDetails struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }
    if err := c.BindJSON(&userDetails); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }
    if err := c.BindJSON(&userDetails); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }
	var tipo string = "normal"
    err := services.RegisterS(userDetails.Username, userDetails.Password, tipo)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Registration successful"})
}
