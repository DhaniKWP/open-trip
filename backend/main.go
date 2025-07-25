package main

import (
	"backend/database"
	"backend/models"
	"backend/routes"
)

func main() {
	database.ConnectDB()
	database.DB.AutoMigrate(&models.Trip{})

	r := routes.SetupRouter()
	r.Run(":8080")
}
