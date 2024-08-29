package route

import (
	handler "ai-project/views/handlers"

	"github.com/gofiber/fiber/v2"
)

func authRoute(app *fiber.App) {
	app.Route("/auth", func(r fiber.Router) {
		r.Post("/sign-in", handler.SignInHandler)
		r.Post("/sign-up", handler.SignUpHandler)

		r.Get("/authenticate/:token", handler.AuthenticateHandler)
		r.Post("/refresh-token/:token", handler.RefreshTokenHandler)
	}, "auth.")

}
