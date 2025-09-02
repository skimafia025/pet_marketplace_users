package routes

import (
	"net/http"
	"pet_marketplace_users/controllers"
	"pet_marketplace_users/logging"
	"pet_marketplace_users/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	logging.SetupLogger()
	r := gin.Default()
	r.Use(logging.RequestLogger())
	r.Use(middlewares.CORSMiddleware())

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "working",
		})
	})

	api := r.Group("/api/v1")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", controllers.Register)
			auth.POST("/login", controllers.Login)
		}
	}

	return r
}
