package cron

import (
	errorLogger "ai-project/utils/error-logger"
	"os"

	"github.com/robfig/cron/v3"
)

var ENV = os.Getenv("APP_ENV")
var ENGINE = os.Getenv("ENGINE")

func Register() {

	c := cron.New()

	go func() {
		if _, err := c.AddFunc("*/1 * * * *", func() { cleanUpUserTokens() }); err != nil {
			errorLogger.CaptureException("cleanUpUserTokens", err)
		}
	}()

	c.Start()

}
