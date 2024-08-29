package route

import (
	handler "ai-project/views/handlers"
	middleware "ai-project/views/middlewares"

	"github.com/gofiber/fiber/v2"
)

func predictionsRoute(app *fiber.App) {
	app.Route("/predictions", func(r fiber.Router) {

		r.Post("/", middleware.AuthMiddleware, handler.GetPredictionsHandler)
		r.Get("/", middleware.AuthMiddleware, handler.GetPredictionsHandler)
		r.Get("/details/:id", middleware.AuthMiddleware, handler.GetSinglePredictionHandler)
		r.Get("/file/download/:id", handler.PredictionDownloadHandler)
		r.Post("/create", middleware.AuthMiddleware, handler.CreatePredictionHandler)

	}, "predictions.")
}
