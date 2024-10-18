package routes

import (
	"go-api/src/http/controllers"

	"github.com/gofiber/fiber/v3"
)

func PublicRoutes(router fiber.Router) {
	router.Get("/", controllers.WelcomeController)
}
