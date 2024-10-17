package controllers

import (
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
)

// Response represents the structure of the JSON response
type Response struct {
	ENVIRONMENT string `json:"ENVIRONMENT"`
	API_NAME    string `json:"API_NAME"`
	API_DESC    string `json:"API_DESC"`
}

func WelcomeController(c fiber.Ctx) error {
	// Loading environment variables
	godotenv.Load()

	// Creates the response structure
	response := Response{
		ENVIRONMENT: os.Getenv("ENVIRONMENT"),
		API_NAME:    os.Getenv("API_NAME"),
		API_DESC:    os.Getenv("API_DESC"),
	}

	// Returns the structure as a JSON response
	return c.Status(fiber.StatusOK).JSON(response)
}
