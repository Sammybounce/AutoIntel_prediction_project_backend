package handler

import (
	"fmt"

	controller "ai-project/controllers"
	model "ai-project/models"
	query "ai-project/utils/queries"
	structValidator "ai-project/utils/struct-validator"

	"github.com/gofiber/fiber/v2"
)

func GetSingleUserHandler(c *fiber.Ctx) error {

	res, code, err := controller.GetSingleUserController(c.Params("id"))
	if err != nil {
		return responseWithErr(c, code, fmt.Sprintf("Err: %v", err.Error()))
	}

	return respondWithJSON(c, code, res)
}

func GetUsersHandler(c *fiber.Ctx) error {

	data := &model.QueryParams{
		Groups: &[]model.FilterGroup{},
	}
	if err := c.BodyParser(data); err != nil {
		return responseWithErr(c, 400, fmt.Sprintf("Err passing to JSON, Something is wrong with your request body json: %v", err.Error()))
	}

	if err := structValidator.Validate(data); err != nil {
		return responseWithErr(c, 400, fmt.Sprintf("Err: %v", err.Error()))
	}

	if err := query.QueryUserFieldValidation(data.OrderBy); err != nil {
		return responseWithErr(c, 400, fmt.Sprintf("Err: %v", err.Error()))
	}

	if err := query.FilterValidation("user", data); err != nil {
		return responseWithErr(c, 400, fmt.Sprintf("Err: %v", err))
	}

	res, code, err := controller.GetUsersController(data)
	if err != nil {
		return responseWithErr(c, code, fmt.Sprintf("Err: %v", err.Error()))
	}

	return respondWithJSON(c, code, res)
}

func UpdateUserHandler(c *fiber.Ctx) error {

	headers := c.GetReqHeaders()

	currentUser := ""

	if len(headers["User-Id"]) > 0 {
		currentUser = headers["User-Id"][0]
	}

	data := &model.UserFull{}
	if err := c.BodyParser(data); err != nil {
		return responseWithErr(c, 400, fmt.Sprintf("Err passing to JSON, Something is wrong with your request body json: %v", err.Error()))
	}

	if err := structValidator.Validate(data); err != nil {
		return responseWithErr(c, 400, fmt.Sprintf("Err: %v", err.Error()))
	}

	res, code, err := controller.UpdateUserController(c.Params("id"), currentUser, data)
	if err != nil {
		return responseWithErr(c, code, fmt.Sprintf("Err: %v", err.Error()))
	}

	return respondWithJSON(c, code, res)
}

func DeleteUserHandler(c *fiber.Ctx) error {

	headers := c.GetReqHeaders()

	currentUser := ""

	if len(headers["User-Id"]) > 0 {
		currentUser = headers["User-Id"][0]
	}

	res, code, err := controller.DeleteUserController(c.Params("id"), currentUser)
	if err != nil {
		return responseWithErr(c, code, fmt.Sprintf("Err: %v", err.Error()))
	}

	return respondWithJSON(c, code, res)
}
