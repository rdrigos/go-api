package main

import (
	"go-api/src/http"

	"github.com/joho/godotenv"
)

func init() {
	// Loading environment variables
	godotenv.Load()
}

func main() {
	// Creates an instance of the server
	server := http.Server()

	// Start the server
	server.Start()
}
