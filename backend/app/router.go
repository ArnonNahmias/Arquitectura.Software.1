package router

import (
	"backend/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Middlewares
	// router.Use(someMiddleware)

	// Rutas y controladores
	router.Use(allowCORS)
	router.GET("/courses", getCourses)
	router.POST("/courses/:id", controllers.SearchbyID)
	router.POST("/login", controllers.Login)
	router.GET("/search?query=", controllers.Search)
	router.GET("/users", getUsers)
	router.POST("/users", createUser)

	return router
}

func allowCORS(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, X-Auth-Token")
	c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Next()
}

// Controladores
func getCourses(c *gin.Context) {
	// Implementación del controlador
	c.JSON(http.StatusOK, gin.H{"message": "getCourses"})
}

func getUsers(c *gin.Context) {
	// Implementación del controlador
	c.JSON(http.StatusOK, gin.H{"message": "getUsers"})
}

func createUser(c *gin.Context) {
	// Implementación del controlador
	c.JSON(http.StatusOK, gin.H{"message": "createUser"})
}
