package handler

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	controller "ai-project/controllers"
	model "ai-project/models"
	Database "ai-project/utils/database"
	errorLogger "ai-project/utils/error-logger"
	query "ai-project/utils/queries"
	structValidator "ai-project/utils/struct-validator"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreatePredictionHandler(c *fiber.Ctx) error {

	m := c.FormValue("model", "decision-tree")
	y := c.FormValue("startYear", "2025")
	y2 := c.FormValue("numberOfYears", "2035")

	if m != "decision-tree-model" && m != "random-forest-model" {
		return responseWithErr(c, 400, fmt.Sprintf("Err: %v", "Model not found"))
	}

	if m == "decision-tree-model" {
		m = "trained-data/decision-tree-trained-model.joblib"
	} else {
		m = "trained-data/random-forest-trained-model.joblib"
	}

	startYear, err := strconv.Atoi(y)
	if err != nil {
		return responseWithErr(c, 400, fmt.Sprintf("Err: %v", err.Error()))
	}

	endYear, err := strconv.Atoi(y2)
	if err != nil {
		return responseWithErr(c, 400, fmt.Sprintf("Err: %v", err.Error()))
	}

	if (startYear+endYear) > 10 || (startYear+endYear) < 1 {
		responseWithErr(c, 400, fmt.Sprintf("Err: %v", "Number of years must be between 1 and 10"))
	}

	file, err := c.FormFile("file")
	if err != nil {
		return responseWithErr(c, 400, fmt.Sprintf("Err: %v", err.Error()))
	}

	if !strings.EqualFold(filepath.Ext(file.Filename), ".csv") {
		return c.Status(fiber.StatusBadRequest).SendString("Only CSV files are allowed")
	}

	id, err := uuid.NewV7()
	if err != nil {
		return responseWithErr(c, 400, fmt.Sprintf("Err: %v", err.Error()))
	}

	filePath := fmt.Sprintf("test-data/%v.csv", id.String())

	if err := c.SaveFile(file, filePath); err != nil {
		return responseWithErr(c, 400, fmt.Sprintf("Err: %v", fmt.Errorf("failed to save file: %v", err.Error())))
	}

	newFile, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer newFile.Close()

	reader := csv.NewReader(newFile)

	records, err := reader.ReadAll()
	if err != nil {
		return responseWithErr(c, 400, fmt.Sprintf("Err: %v", err.Error()))
	}

	for i, record := range records {
		if i != 0 {
			break
		}

		brand := record[0]
		model := record[1]
		year := record[2]
		price := record[3]
		transmission := record[4]
		mileage := record[5]
		tax := record[6]
		mpg := record[7]
		fuelType := record[8]
		engineSize := record[9]

		if strings.ToLower(brand) != "brand" {
			return responseWithErr(c, 400, fmt.Sprintf("Err: %v", "Invalid CSV file brand is missing"))
		}

		if strings.ToLower(model) != "model" {
			return responseWithErr(c, 400, fmt.Sprintf("Err: %v", "Invalid CSV file model is missing"))
		}

		if strings.ToLower(year) != "year" {
			return responseWithErr(c, 400, fmt.Sprintf("Err: %v", "Invalid CSV file year is missing"))
		}

		if strings.ToLower(price) != "price" {
			return responseWithErr(c, 400, fmt.Sprintf("Err: %v", "Invalid CSV file price is missing"))
		}

		if strings.ToLower(transmission) != "transmission" {
			return responseWithErr(c, 400, fmt.Sprintf("Err: %v", "Invalid CSV file transmission is missing"))
		}

		if strings.ToLower(mileage) != "mileage" {
			return responseWithErr(c, 400, fmt.Sprintf("Err: %v", "Invalid CSV file mileage is missing"))
		}

		if strings.ToLower(tax) != "tax" {
			return responseWithErr(c, 400, fmt.Sprintf("Err: %v", "Invalid CSV file tax is missing"))
		}

		if strings.ToLower(mpg) != "mpg" {
			return responseWithErr(c, 400, fmt.Sprintf("Err: %v", "Invalid CSV file mpg is missing"))
		}

		if strings.ToLower(fuelType) != "fueltype" {
			return responseWithErr(c, 400, fmt.Sprintf("Err: %v", "Invalid CSV file fuelType is missing"))
		}

		if strings.ToLower(engineSize) != "enginesize" {
			return responseWithErr(c, 400, fmt.Sprintf("Err: %v", "Invalid CSV file engineSize is missing"))
		}
	}

	data := &model.PredictionCreate{
		Id:        id.String(),
		Model:     m,
		FilePath:  fmt.Sprintf("test-data/%v.csv", id.String()),
		StartYear: startYear,
		EndYear:   endYear,
	}

	if err := structValidator.Validate(data); err != nil {
		return responseWithErr(c, 400, fmt.Sprintf("Err: %v", err.Error()))
	}

	res, code, err := controller.CreatePredictionController(data)
	if err != nil {
		return responseWithErr(c, code, fmt.Sprintf("Err: %v", err.Error()))
	}

	return respondWithJSON(c, code, res)
}

func GetSinglePredictionHandler(c *fiber.Ctx) error {

	res, code, err := controller.GetSinglePredictionController(c.Params("id"))
	if err != nil {
		return responseWithErr(c, code, fmt.Sprintf("Err: %v", err.Error()))
	}

	return respondWithJSON(c, code, res)
}

func PredictionDownloadHandler(c *fiber.Ctx) error {

	db := Database.Connect()
	defer db.Close()

	sqlStatement := fmt.Sprintf(`
		SELECT 

			P.group_id, P.brand, P.model, P.year, P.future_year, P.prediction_model, P.predicted_price, P.created_at


		FROM 
			predictions AS P

		WHERE
			P.group_id = '%v' 
			AND P.deleted = false

		ORDER BY
			P.future_year ASC;

	`,
		c.Params("id"),
	)

	rows, err := db.Query(sqlStatement)
	if err != nil {
		errorLogger.CaptureException("GetSinglePredictionController--1", fmt.Errorf("%v | error: %v", sqlStatement, err))

		return responseWithErr(c, 500, fmt.Sprintf("Err: %v", err.Error()))
	}
	defer rows.Close()

	prediction := []model.Prediction{}

	for rows.Next() {
		p := model.PredictionSQL{}

		if err := rows.Scan(
			&p.Id,
			&p.Brand,
			&p.Model,
			&p.Year,
			&p.FutureYear,
			&p.PredictionModel,
			&p.PredictedPrice,
			&p.CreatedAt,
		); err != nil {
			errorLogger.CaptureException("GetSinglePredictionController--2", err)

			return responseWithErr(c, 500, fmt.Sprintf("Err: %v", err.Error()))
		}

		prediction = append(prediction, model.Prediction{
			Id:              p.Id.String,
			Brand:           p.Brand.String,
			Model:           p.Model.String,
			Year:            p.Year.Int64,
			FutureYear:      p.FutureYear.Int64,
			PredictionModel: p.PredictionModel.String,
			PredictedPrice:  p.PredictedPrice.String,
			CreatedAt:       p.CreatedAt.Time.Format("2006-01-02T15:04:05-07:00"),
		})
	}

	_, err = os.ReadFile("predictions/" + c.Params("id") + ".csv")
	if err != nil {
		data := [][]string{
			{"brand", "model", "year", "future_year", "predicted_price"},
		}

		for _, v := range prediction {
			data = append(data, []string{v.Brand, v.Model, string(v.Year), string(v.FutureYear), v.PredictedPrice})
		}

		file, err := os.Create("predictions/" + c.Params("id") + ".csv")
		if err != nil {
			responseWithErr(c, 500, fmt.Sprintf("Failed to create file: %v", err.Error()))
		}
		defer file.Close()

		writer := csv.NewWriter(file)
		defer writer.Flush()

		for _, record := range data {
			if err := writer.Write(record); err != nil {
				return responseWithErr(c, 500, fmt.Sprintf("Failed to write record to file: %v", err.Error()))
			}
		}
	}

	fmt.Println("here")

	c.Set("Content-Type", "text/csv")

	return c.SendFile("predictions/"+c.Params("id")+".csv", true)
}

func GetPredictionsHandler(c *fiber.Ctx) error {

	data := &model.QueryParams{
		Groups: &[]model.FilterGroup{},
	}
	if err := c.BodyParser(data); err != nil {
		return responseWithErr(c, 400, fmt.Sprintf("Err passing to JSON, Something is wrong with your request body json: %v", err.Error()))
	}

	if err := structValidator.Validate(data); err != nil {
		return responseWithErr(c, 400, fmt.Sprintf("Err: %v", err.Error()))
	}

	if err := query.QueryPredictionFieldValidation(data.OrderBy); err != nil {
		return responseWithErr(c, 400, fmt.Sprintf("Err: %v", err.Error()))
	}

	if err := query.FilterValidation("prediction", data); err != nil {
		return responseWithErr(c, 400, fmt.Sprintf("Err: %v", err))
	}

	res, code, err := controller.GetPredictionsController(data)
	if err != nil {
		return responseWithErr(c, code, fmt.Sprintf("Err: %v", err.Error()))
	}

	return respondWithJSON(c, code, res)
}
