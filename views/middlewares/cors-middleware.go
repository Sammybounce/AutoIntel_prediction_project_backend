package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func corsMiddleware(r *fiber.App) {
	r.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS",
		AllowHeaders:     "*",
		ExposeHeaders:    "x-auth-token",
		AllowCredentials: false,
	}))
}
