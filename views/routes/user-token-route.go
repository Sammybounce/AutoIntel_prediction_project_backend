package route

import (
	handler "ai-project/views/handlers"

	"github.com/gofiber/fiber/v2"
)

func userTokenRoute(r fiber.Router) {
	r.Route("/token", func(r fiber.Router) {
		r.Post("/", handler.GetUsersTokenHandler)
		r.Get("/", handler.GetUsersTokenHandler)
		r.Get("/details/:userId", handler.GetSingleUserTokenHandler)
	}, "token.")
}
