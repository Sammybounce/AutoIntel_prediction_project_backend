package middleware

import (
	"fmt"

	Database "ai-project/utils/database"
	errorLogger "ai-project/utils/error-logger"

	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(c *fiber.Ctx) error {

	headers := c.GetReqHeaders()

	if len(headers["X-Auth-Token"]) > 0 {

		token := headers["X-Auth-Token"][0]

		db := Database.Connect()

		var userId, exp string
		sqlStatement := fmt.Sprintf(`SELECT user_id, expire_at FROM user_tokens WHERE token = '%v' AND expire_at > CURRENT_TIMESTAMP;`, token)
		rows, err := db.Query(sqlStatement)

		if err != nil {
			errorLogger.CaptureException("AuthMiddleware--1", fmt.Errorf("%v | error: %v", sqlStatement, err))

			return c.Status(500).JSON(err)
		}

		defer rows.Close()

		for rows.Next() {
			rows.Scan(&userId, &exp)
		}

		if exp == "" {
			return c.Status(400).SendString("Invalid token provide in header x-auth-token or token has expired")
		}

		c.Request().Header.Set("User-Id", userId)
	}

	return c.Next()
}
