package view

import (
	"fmt"
	"os"

	controller "ai-project/controllers"
	"ai-project/cron-jobs"
	model "ai-project/models"
	middleware "ai-project/views/middlewares"
	route "ai-project/views/routes"

	"github.com/gofiber/fiber/v2"
)

func CreateNewServer() *model.Server {
	s := &model.Server{}
	s.Router = fiber.New(fiber.Config{
		BodyLimit: 1000 * 1024 * 1024,
	})

	return s
}

func StartServer() {

	go cron.Register()

	controller.PresetDefaults()

	r := CreateNewServer()

	r.Middleware(middleware.MountMiddleware)
	r.Routes(route.MountRoutes)

	APP_PORT := os.Getenv("APP_PORT")

	r.Router.Listen(fmt.Sprintf(":%v", APP_PORT))
}
