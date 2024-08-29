package cache

import (
	"encoding/json"
	"os"

	model "ai-project/models"
)

func VehicleCache() *[]model.Vehicle {

	v := []model.Vehicle{}

	if jsonData, err := os.ReadFile("cache/files/vehicle.json"); err == nil {
		json.Unmarshal(jsonData, &v)
	}

	return &v
}
