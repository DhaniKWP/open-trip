package main

import (
	"log"
	"opentrip-backend/config"
	"opentrip-backend/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.InitDB()

	r := gin.Default()
	routes.SetupRoutes(r)

	r.Run(":8080")
}
