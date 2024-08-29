package errorLogger

import (
	"encoding/json"
	"os"

	model "ai-project/models"
)

func ErrorCache() *[]model.CaptureError {

	e := []model.CaptureError{}

	if jsonData, err := os.ReadFile("cache/files/error.json"); err == nil {
		json.Unmarshal(jsonData, &e)
	}

	return &e
}
