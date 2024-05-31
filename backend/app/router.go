package main

import (
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.Default()

	// Middlewares
	// router.Use(someMiddleware)

	// Rutas y controladores
	router.GET("/courses", getCourses)
	router.POST("/login", login)
	router.GET("/search", search)
	router.GET("/users", getUsers)
	router.POST("/users", createUser)

	return router
}

func main() {
	router := setupRouter()
	router.Run(":8080")
}
