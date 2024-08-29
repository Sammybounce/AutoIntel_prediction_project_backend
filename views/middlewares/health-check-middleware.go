package middleware

import (
	Database "ai-project/utils/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
)

func healthCheckMiddleware(app *fiber.App) {
	app.Use(healthcheck.New(healthcheck.Config{
		LivenessProbe: func(c *fiber.Ctx) bool {
			return true
		},
		LivenessEndpoint: "/ping",
		ReadinessProbe: func(c *fiber.Ctx) bool {
			err := Database.Connect().Ping()
			return err == nil
		},
		ReadinessEndpoint: "/ready",
	}))
}
