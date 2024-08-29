package model

import "database/sql"

type PredictionCreate struct {
	Id        string `json:"id" validate:"required"`
	Model     string `json:"model" validate:"required"`
	FilePath  string `json:"filePath" validate:"required"`
	StartYear int    `json:"startYear" validate:"required"`
	EndYear   int    `json:"endYear" validate:"required"`
}

type Prediction struct {
	Id              string `json:"id"`
	Brand           string `json:"brand"`
	Model           string `json:"model"`
	Year            int64  `json:"year"`
	FutureYear      int64  `json:"futureYear"`
	PredictionModel string `json:"predictionModel"`
	PredictedPrice  string `json:"predictedPrice"`
	CreatedAt       string `json:"createdAt"`
}

type PredictionSQL struct {
	Id, Brand, Model, PredictionModel, PredictedPrice sql.NullString
	Year, FutureYear                                  sql.NullInt64
	CreatedAt                                         sql.NullTime
}
