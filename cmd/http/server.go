package http

import (
	"fmt"
	"log"
	"os"

	"go-api/cmd/http/middlewares"

	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
)

// Loads the necessary configurations and packages on initialization
func init() {
	// Loading environment variables
	godotenv.Load()
}

type ServerType struct {
	server *fiber.App
}

func Server() *ServerType {
	server := fiber.New(fiber.Config{
		AppName: os.Getenv("API_NAME"),
	})

	// Configure middlewares
	middlewares.SetupCors(server)
	middlewares.SetupHelmet(server)

	// Configures the server routes
	Router(server.Group("/api/"))

	return &ServerType{server: server}
}

func (s *ServerType) Start() {
	// Gets the host address and port from the environment variables API_HOST and API_PORT
	host := os.Getenv("API_HOST")
	port := os.Getenv("API_PORT")

	// Starts the HTTP server and listens on the specified address
	if err := s.server.Listen(fmt.Sprintf("%s:%s", host, port)); err != nil {
		log.Fatal(err)
	}
}
