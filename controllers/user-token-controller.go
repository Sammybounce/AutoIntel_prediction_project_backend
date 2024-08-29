package controller

import (
	"fmt"

	model "ai-project/models"
	Database "ai-project/utils/database"
	errorLogger "ai-project/utils/error-logger"
)

func GetUsersTokenController() (r *model.ResponseWithMultipleData, code int, err error) {

	db := Database.Connect()
	defer db.Close()

	selectStatement := `
		SELECT 
			id, 
			user_id, 
			token, 
			expire_at 
			
		FROM 
			user_tokens;
	`

	rows, err := db.Query(selectStatement)
	if err != nil {
		errorLogger.CaptureException("GetUsersTokenController--1", fmt.Errorf("%v | error: %v", selectStatement, err))

		return nil, 500, err
	}

	defer rows.Close()

	userTokens := []model.UserToken{}

	for rows.Next() {
		ut := model.UserTokenSQL{}

		if err := rows.Scan(&ut.Id, &ut.UserId, &ut.Token, &ut.ExpireAt); err != nil {
			errorLogger.CaptureException("GetUsersTokenController--2", err)

			return nil, 500, err
		}

		if ut.Id.Valid {
			userTokens = append(userTokens, model.UserToken{
				Id:       ut.Id.String,
				UserId:   ut.UserId.String,
				Token:    ut.Token.String,
				ExpireAt: ut.ExpireAt.Time.Format("2006-01-02T15:04:05-07:00"),
			})
		}
	}

	var res = &model.ResponseWithMultipleData{
		Data:         userTokens,
		Type:         "success",
		Status:       200,
		Message:      "success",
		PageNumber:   1,
		BatchNumber:  len(userTokens),
		TotalRecords: len(userTokens),
	}

	return res, 200, nil
}

func GetSingleUserTokenController(id string) (r *model.ResponseWithMultipleData, code int, err error) {
	db := Database.Connect()
	defer db.Close()

	selectStatement := fmt.Sprintf(`
		SELECT 
			id, 
			user_id, 
			token, 
			expire_at 
			
		FROM 
			user_tokens 
			
		WHERE 
			user_id = '%v';
	`,
		id,
	)

	rows, err := db.Query(selectStatement)
	if err != nil {
		errorLogger.CaptureException("GetSingleUserTokenController--1", fmt.Errorf("%v | error: %v", selectStatement, err))

		return nil, 500, err
	}

	defer rows.Close()

	userTokens := []model.UserToken{}

	for rows.Next() {
		ut := model.UserTokenSQL{}

		if err := rows.Scan(&ut.Id, &ut.UserId, &ut.Token, &ut.ExpireAt); err != nil {
			errorLogger.CaptureException("GetSingleUserTokenController--2", err)

			return nil, 500, err
		}

		if ut.Id.Valid {
			userTokens = append(userTokens, model.UserToken{
				Id:       ut.Id.String,
				UserId:   ut.UserId.String,
				Token:    ut.Token.String,
				ExpireAt: ut.ExpireAt.Time.Format("2006-01-02T15:04:05-07:00"),
			})
		}
	}

	var res = &model.ResponseWithMultipleData{
		Data:         userTokens,
		Type:         "success",
		Status:       200,
		Message:      "success",
		PageNumber:   1,
		BatchNumber:  len(userTokens),
		TotalRecords: len(userTokens),
	}

	return res, 200, nil
}
