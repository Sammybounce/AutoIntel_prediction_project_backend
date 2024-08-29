package route

import (
	handler "ai-project/views/handlers"
	middleware "ai-project/views/middlewares"

	"github.com/gofiber/fiber/v2"
)

func userRoute(app *fiber.App) {
	app.Route("/users", func(r fiber.Router) {
		r.Use(middleware.AuthMiddleware)

		r.Post("/", handler.GetUsersHandler)
		r.Get("/", handler.GetUsersHandler)
		r.Get("/details/:id", handler.GetSingleUserHandler)
		// r.Post("/create", handler.CreateUserHandler)
		r.Put("/update/:id", handler.UpdateUserHandler)
		r.Post("/delete/:id", handler.DeleteUserHandler)

		userTokenRoute(r)

	}, "user.")
}
