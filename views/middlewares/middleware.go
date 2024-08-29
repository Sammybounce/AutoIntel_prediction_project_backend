package middleware

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

var ENV string = os.Getenv("APP_ENV")

func MountMiddleware(app *fiber.App) {

	loggerMiddleware(app)
	corsMiddleware(app)
	healthCheckMiddleware(app)
	compressMiddleware(app)
	helmetMiddleware(app)

}
