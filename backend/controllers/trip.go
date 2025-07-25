package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"opentrip-backend/database"
	"opentrip-backend/models"
)

func GetTrips(c *gin.Context) {
	var trips []models.Trip
	database.DB.Find(&trips)
	c.JSON(http.StatusOK, trips)
}

func CreateTrip(c *gin.Context) {
	var input models.Trip
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&input)
	c.JSON(http.StatusOK, input)
}
