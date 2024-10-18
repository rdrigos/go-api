package middlewares

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func SetupCors(server *fiber.App) {
	server.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET, POST, PUT, DELETE, OPTIONS"},
		AllowHeaders: []string{"Content-Type, Authorization"},
	}))
}
