package main

import (
	"log"
	"net/http"
	"to-do-list/internal/app"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Printf("error loading .env file: %v", err)
		return
	}

	app.SetupRouter()
	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}
