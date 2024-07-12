package app

import (
	"backend/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Middlewares
	router.Use(allowCORS)

	// Rutas y controladores
	router.GET("/courses", controllers.GetCourses)
	router.GET("/courses/:id", controllers.SearchByID)
	router.GET("/courses/name/:name", controllers.SearchByName)	
	router.GET("/courses/category/:category", controllers.GetCoursesByCategory)
	// router.POST("/courses", controllers.CreateCourse)
	// router.DELETE("/courses/:id", controllers.DeleteCourse)
	router.POST("/login", controllers.Login)
	router.GET("/subscriptions", controllers.GetSubscriptions)
	router.POST("/subscriptions", controllers.CreateSubscription)
	router.DELETE("/subscriptions/:id", controllers.DeleteSubscription)
	router.POST("/register", controllers.RegisterC)

	// Rutas protegidas
	protected := router.Group("/protected")
	protected.Use(Auth()) // Usamos el middleware de autenticación
	{
		protected.GET("/", controllers.ProtectedEndpoint)
	}
	return router
}

func allowCORS(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, X-Auth-Token")
	c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
		return
	}
	c.Next()
}
