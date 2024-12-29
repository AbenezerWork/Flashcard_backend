package main

import (
	"flashcard/config"
	"flashcard/models"
	"flashcard/routes"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	// Initialize the database
	config.InitDB()

	// Auto migrate the models
	config.DB.AutoMigrate(&models.User{}, &models.Flashcard{}, &models.Deck{})

	// Set up the router
	r := routes.SetupRouter(config.DB)

	// Run the server
	r.Run(":8080")
}
