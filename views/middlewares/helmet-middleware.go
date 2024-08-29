package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"
)

func helmetMiddleware(app *fiber.App) {
	app.Use(helmet.New())
}
