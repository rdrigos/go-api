package http

import (
	"go-api/src/http/routes"

	"github.com/gofiber/fiber/v3"
)

func Router(router fiber.Router) {
	routes.PublicRoutes(router)
}
