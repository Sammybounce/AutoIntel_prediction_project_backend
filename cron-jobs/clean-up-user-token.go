package cron

import (
	errorLogger "ai-project/utils/error-logger"
	"fmt"

	Database "ai-project/utils/database"
)

func cleanUpUserTokens() {
	db := Database.Connect()
	defer db.Close()

	deleteStatement := `DELETE FROM user_tokens WHERE expire_at < NOW();`
	if _, err := db.Exec(deleteStatement); err != nil {
		errorLogger.CaptureException("cleanUpUserTokens", fmt.Errorf("%v | error: %v", deleteStatement, err))
	}
}
