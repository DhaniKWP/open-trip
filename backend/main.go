package main

import (
    "opentrip-backend/database"
    "opentrip-backend/models"
    "opentrip-backend/routes"
)

func main() {
	database.ConnectDB()
	database.DB.AutoMigrate(&models.Trip{})

	r := routes.SetupRouter()
	r.Run(":8080")
}
