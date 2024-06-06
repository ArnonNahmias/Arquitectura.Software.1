package controllers

import (
	"backend/domain"
	"backend/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	var request domain.LoginRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	token, err := services.Login(request.Username, request.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, err.Error())
		return
	}

	c.JSON(http.StatusOK, domain.LoginResponse{
		Token: token,
	})
}
