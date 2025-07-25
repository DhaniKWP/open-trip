package main

import (
	"backend/database"
	"backend/routes"
)

func main() {
	database.ConnectDB()
	r := routes.SetupRouter()
	r.Run(":8080")
}

