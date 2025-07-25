package routes

import (
	"github.com/gin-gonic/gin"
	"opentrip-backend/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		api.GET("/trips", controllers.GetTrips)
		api.POST("/trips", controllers.CreateTrip)
	}

	return r
}
