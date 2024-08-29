package handler

import (
	"fmt"

	controller "ai-project/controllers"
	model "ai-project/models"
	structValidator "ai-project/utils/struct-validator"

	"github.com/gofiber/fiber/v2"
)

func SignInHandler(c *fiber.Ctx) error {

	data := &model.SignIn{}
	if err := c.BodyParser(data); err != nil {
		return responseWithErr(c, 400, fmt.Sprintf("Err passing to JSON, Something is wrong with your request body json: %v", err.Error()))
	}

	if err := structValidator.Validate(data); err != nil {
		return responseWithErr(c, 400, fmt.Sprintf("Err: %v", err.Error()))
	}

	res, code, err := controller.SignInController(data)
	if err != nil {
		return responseWithErr(c, code, fmt.Sprintf("Err: %v", err.Error()))
	}

	return respondWithJSON(c, code, res)
}

func SignUpHandler(c *fiber.Ctx) error {

	data := &model.SignUp{}
	if err := c.BodyParser(data); err != nil {
		return responseWithErr(c, 400, fmt.Sprintf("Err passing to JSON, Something is wrong with your request body json: %v", err.Error()))
	}

	if err := structValidator.Validate(data); err != nil {
		return responseWithErr(c, 400, fmt.Sprintf("Err: %v", err.Error()))
	}

	res, code, err := controller.SignUpController(data)
	if err != nil {
		return responseWithErr(c, code, fmt.Sprintf("Err: %v", err.Error()))
	}

	return respondWithJSON(c, code, res)
}

func AuthenticateHandler(c *fiber.Ctx) error {

	res, code, err := controller.AuthenticateController(c.Params("token"))
	if err != nil {
		return responseWithErr(c, code, fmt.Sprintf("Err: %v", err.Error()))
	}

	return respondWithJSON(c, code, res)
}

func RefreshTokenHandler(c *fiber.Ctx) error {

	res, code, err := controller.RefreshTokenController(c.Params("token"))
	if err != nil {
		return responseWithErr(c, code, fmt.Sprintf("Err: %v", err.Error()))
	}

	return respondWithJSON(c, code, res)
}
