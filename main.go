package main

import (
	"go-api/cmd/http"
)

func main() {
	// Creates an instance of the server
	server := http.Server()

	// Start the server
	server.Start()
}
