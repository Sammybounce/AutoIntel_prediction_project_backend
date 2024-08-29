package handler

import (
	"fmt"

	controller "ai-project/controllers"

	"github.com/gofiber/fiber/v2"
)

func GetUsersTokenHandler(c *fiber.Ctx) error {

	res, code, err := controller.GetUsersTokenController()
	if err != nil {
		return responseWithErr(c, code, fmt.Sprintf("Err: %v", err.Error()))
	}

	return respondWithJSON(c, code, res)
}

func GetSingleUserTokenHandler(c *fiber.Ctx) error {

	res, code, err := controller.GetSingleUserTokenController(c.Params("userId"))
	if err != nil {
		return responseWithErr(c, code, fmt.Sprintf("Err: %v", err.Error()))
	}

	return respondWithJSON(c, code, res)
}
