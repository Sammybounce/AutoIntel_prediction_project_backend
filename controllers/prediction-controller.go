package controller

import (
	model "ai-project/models"
	Database "ai-project/utils/database"
	errorLogger "ai-project/utils/error-logger"
	query "ai-project/utils/queries"
	"encoding/csv"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/google/uuid"
)

func CreatePredictionController(data *model.PredictionCreate) (r *model.ResponseWithSingleData, code int, err error) {

	db := Database.Connect()
	defer db.Close()

	file, err := os.ReadFile("notebooks/price-predict.ipynb")
	if err != nil {
		return nil, 500, err
	}

	predictionFilePath := fmt.Sprintf("predictions/%v.csv", data.Id)

	tempNotebook := strings.Replace(string(file), "trained-data/decision-tree-trained-model.joblib", fmt.Sprintf("../%v", data.Model), 1)
	tempNotebook = strings.Replace(tempNotebook, "sample-data/dataset.csv", fmt.Sprintf("../%v", data.FilePath), 1)
	tempNotebook = strings.Replace(tempNotebook, "predictions/decision-tree-model-training.csv", fmt.Sprintf("../%v", predictionFilePath), 1)
	tempNotebook = strings.Replace(tempNotebook, "future_years = range(2025, 2035)", fmt.Sprintf("future_years = range(%v, %v)", data.StartYear, data.EndYear), 1)

	err = os.WriteFile(fmt.Sprintf("notebooks/temp/%v.ipynb", data.Id), []byte(tempNotebook), 0644)
	if err != nil {
		return nil, 500, err
	}

	var predict bool

	cmdStr := fmt.Sprintf("jupyter nbconvert --to notebook --execute notebooks/temp/%v.ipynb --stdout", data.Id)

	cmd := exec.Command(strings.Split(cmdStr, " ")[0], strings.Split(cmdStr, " ")[1:]...)

	stdout, err := cmd.Output()
	if err != nil {
		errorLogger.CaptureException("CreatePredictionController--1ab", err)

		return nil, 500, fmt.Errorf("failed to execute command: %s", err)
	}

	if strings.Contains(string(stdout), "success") {
		predict = true
	}

	if !predict {
		os.Remove(fmt.Sprintf("notebooks/temp/%v.ipynb", data.Id))

		return nil, 500, fmt.Errorf("failed to predict")
	}

	os.Remove(fmt.Sprintf("notebooks/temp/%v.ipynb", data.Id))

	newFile, err := os.Open(predictionFilePath)
	if err != nil {
		return nil, 500, err
	}
	defer newFile.Close()

	reader := csv.NewReader(newFile)

	records, err := reader.ReadAll()
	if err != nil {
		return nil, 500, err
	}

	var execStatement, m string

	if data.Model == "trained-data/decision-tree-trained-model.joblib" {
		m = "decision tree"
	} else {
		m = "random forest"
	}

	for i, record := range records {
		if i == 0 {
			continue
		}

		id, err := uuid.NewV7()
		if err != nil {
			return nil, 500, err
		}

		brand := record[0]
		model := record[1]
		year := record[2]
		futureYear := record[3]
		predictedPrice := record[4]

		execStatement = fmt.Sprintf(`
			%v

			INSERT INTO predictions 
				(id, group_id, brand, model, year, future_year, prediction_model, predicted_price) 
			VALUES 
				('%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v');
		`,
			execStatement,
			id,
			data.Id,
			brand,
			model,
			year,
			futureYear,
			m,
			predictedPrice,
		)
	}

	tx, err := db.Begin()
	if err != nil {
		return nil, 500, err
	}

	_, err = tx.Exec(execStatement)
	if err != nil {

		errorLogger.CaptureException("CreatePredictionController--1", err)

		if err = tx.Rollback(); err != nil {
			return nil, 500, err
		}

		return nil, 500, err
	}

	if err = tx.Commit(); err != nil {
		return nil, 500, err
	}

	res := &model.ResponseWithSingleData{
		Data:    data.Id,
		Type:    "success",
		Message: "Prediction created successfully",
		Status:  200,
	}

	return res, 200, nil
}

func GetPredictionsController(data *model.QueryParams) (r *model.ResponseWithMultipleData, code int, err error) {

	db := Database.Connect()
	defer db.Close()

	from := `
		FROM 
			predictions AS P
	`

	var sqlStatement = fmt.Sprintf(`
		WITH AGGREGATED_DATA AS (

			SELECT COUNT(DISTINCT P.id) AS total_records
			
			%v
				
			--WHERE
		)

		SELECT

			P.group_id, P.brand, P.model, P.year, P.future_year, P.prediction_model, P.predicted_price, P.created_at,

			AD.total_records

		%v

		CROSS JOIN AGGREGATED_DATA AS AD

		--WHERE

		--ORDER_BY
			
		--OFFSET

		--LIMIT

	`,
		from,
		from,
	)

	sqlStatement = fmt.Sprintf("%v", *query.GenerateSQL(data, sqlStatement, "predictions"))

	rows, err := db.Query(sqlStatement)
	if err != nil {
		errorLogger.CaptureException("GetPredictionsController--1", fmt.Errorf("%v | error: %v", sqlStatement, err))

		return nil, 500, err
	}
	defer rows.Close()

	predictions := []model.Prediction{}
	totalRecords := 0

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

			&totalRecords,
		); err != nil {
			errorLogger.CaptureException("GetPredictionsController--2", err)

			return nil, 500, err
		}

		if p.Id.Valid {
			predictions = append(predictions, model.Prediction{
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
	}

	var res = &model.ResponseWithMultipleData{
		Data:         predictions,
		Type:         "success",
		Status:       200,
		Message:      "success",
		PageNumber:   data.PageNumber,
		BatchNumber:  data.BatchNumber,
		TotalRecords: totalRecords,
	}

	return res, 200, nil
}

func GetSinglePredictionController(id string) (r *model.ResponseWithSingleData, code int, err error) {

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
		id,
	)

	rows, err := db.Query(sqlStatement)
	if err != nil {
		errorLogger.CaptureException("GetSinglePredictionController--1", fmt.Errorf("%v | error: %v", sqlStatement, err))

		return nil, 500, err
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

			return nil, 500, err
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

	if len(prediction) == 0 {
		return nil, 400, fmt.Errorf("prediction %v not found", id)
	}

	var res = &model.ResponseWithSingleData{
		Data:    prediction,
		Type:    "success",
		Status:  200,
		Message: "success",
	}

	return res, 200, nil
}
