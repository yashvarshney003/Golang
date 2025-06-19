package main

import (
	"log"
	"golang_crud/db"
	"golang_crud/routes"
)

func main() {
	// Load environment variables into AppConfig
	db.LoadEnv()
	db.ConnectDatabase()
	router := routes.SetupRouter()
    log.Println("🚀 Server running at http://localhost:8080")
    router.Run(":8080")
	// You can now use `connection` to run DB operations.
}
