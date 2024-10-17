package middlewares

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/helmet"
)

func SetupHelmet(server *fiber.App) {
	server.Use(helmet.New())
}
