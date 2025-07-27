package routes

import (
	"opentrip-backend/controllers"
	"opentrip-backend/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")

	api.POST("/register", controllers.Register)
	api.POST("/login", controllers.Login)

	auth := api.Group("/")
	auth.Use(middleware.AuthMiddleware())
	auth.GET("/me", controllers.Me)

}
