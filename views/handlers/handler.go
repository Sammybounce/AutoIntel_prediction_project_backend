package handler

import (
	"log"

	model "ai-project/models"

	"github.com/gofiber/fiber/v2"
)

func responseWithErr(c *fiber.Ctx, code int, msg string) error {
	if code > 499 {
		log.Println("Responding with 5XX error:", msg)
	}

	return respondWithJSON(c, code, &model.ResponseWithSingleData{
		Data:    &struct{}{},
		Status:  code,
		Type:    "error",
		Message: msg,
	})
}

func respondWithJSON(c *fiber.Ctx, code int, payload interface{}) error {
	return c.Status(code).JSON(payload)
}
