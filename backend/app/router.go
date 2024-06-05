package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter() *gin.Engine {
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

// Controladores
func getCourses(c *gin.Context) {
	// Implementación del controlador
	c.JSON(http.StatusOK, gin.H{"message": "getCourses"})
}

func login(c *gin.Context) {
	// Implementación del controlador
	c.JSON(http.StatusOK, gin.H{"message": "login"})
}

func search(c *gin.Context) {
	// Implementación del controlador
	c.JSON(http.StatusOK, gin.H{"message": "search"})
}

func getUsers(c *gin.Context) {
	// Implementación del controlador
	c.JSON(http.StatusOK, gin.H{"message": "getUsers"})
}

func createUser(c *gin.Context) {
	// Implementación del controlador
	c.JSON(http.StatusOK, gin.H{"message": "createUser"})
}
