package query

import (
	"errors"
	"fmt"

	"ai-project/utils/array"
)

var AllowedPredictionFields = []string{
	"prediction.id",
	"prediction.groupId",
	"prediction.brand",
	"prediction.model",
	"prediction.year",
	"prediction.futureYear",
	"prediction.predictionModel",
	"prediction.predictedPrice",
	"prediction.createdAt",
	"prediction.updatedAt",
	"prediction.deleted",
}

func QueryPredictionFieldValidation(field string) error {

	_, check, _ := array.Find[string](&AllowedPredictionFields, func(d string) bool {
		return d == field
	})

	if !check {
		err := fmt.Sprintf("%v is not allowed go to /allowed-query/fields/predictions to see the list of allowed query fields", field)
		return errors.New(err)
	}

	return nil
}
