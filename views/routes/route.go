package route

import (
	model "ai-project/models"
	query "ai-project/utils/queries"

	"github.com/gofiber/fiber/v2"
)

func MountRoutes(app *fiber.App) {

	app.Static("/", "./docs")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile("./docs/index.html")
	})

	app.Route("/allowed-query", func(r fiber.Router) {
		r.Route("/fields", func(r fiber.Router) {
			r.Get("/users", func(c *fiber.Ctx) error {
				data := model.ResponseWithSingleData{
					Data:    query.AllowedUserFields,
					Status:  200,
					Type:    "success",
					Message: "success",
				}

				return c.Status(200).JSON(data)
			})
			r.Get("/predictions", func(c *fiber.Ctx) error {
				data := model.ResponseWithSingleData{
					Data:    query.AllowedPredictionFields,
					Status:  200,
					Type:    "success",
					Message: "success",
				}

				return c.Status(200).JSON(data)
			})
		})

		r.Get("/search-condition", func(c *fiber.Ctx) error {
			data := model.ResponseWithSingleData{
				Data:    query.AllowedSearchCondition,
				Status:  200,
				Type:    "success",
				Message: "success",
			}

			return c.Status(200).JSON(data)
		})
	})

	authRoute(app)
	userRoute(app)
	predictionsRoute(app)
}
