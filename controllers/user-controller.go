// Package controller provides the implementation of various controllers
// responsible for handling business logic related to the user account.

package controller

import (
	"errors"
	"fmt"

	model "ai-project/models"
	Database "ai-project/utils/database"
	errorLogger "ai-project/utils/error-logger"
	query "ai-project/utils/queries"
)

// GetUsersController retrieves user data based on the specified query parameters.
func GetUsersController(data *model.QueryParams) (r *model.ResponseWithMultipleData, code int, err error) {

	db := Database.Connect()
	defer db.Close()

	from := `
		FROM 
			users AS U
	`

	// SQL statement to fetch user data along with profile, address, and phone numbers
	var sqlStatement = fmt.Sprintf(`
		WITH AGGREGATED_DATA AS (

			SELECT COUNT(DISTINCT U.id) AS total_records
			
			%v
				
			--WHERE
		)

		SELECT

			U.id, U.first_name, U.last_name, U.email, U.created_at,

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

	// Generate SQL statement based on query parameters
	sqlStatement = fmt.Sprintf("%v", *query.GenerateSQL(data, sqlStatement, "users"))

	// Execute the SQL query
	rows, err := db.Query(sqlStatement)
	if err != nil {
		errorLogger.CaptureException("GetUsersController--1", fmt.Errorf("%v | error: %v", sqlStatement, err))

		return nil, 500, err
	}
	defer rows.Close()

	// Initialize variables to store user-related data
	users := []model.User{}
	totalRecords := 0

	// Iterate through query results
	for rows.Next() {

		u := model.UserSQL{}

		// Scan data from the database into corresponding structs
		if err := rows.Scan(
			&u.Id, &u.FirstName, &u.LastName, &u.Email, &u.CreatedAt,

			&totalRecords,
		); err != nil {
			errorLogger.CaptureException("GetUsersController--2", err)

			return nil, 500, err
		}

		if u.Id.Valid {
			users = append(users, model.User{
				Id:        u.Id.String,
				FirstName: u.FirstName.String,
				LastName:  u.LastName.String,
				Email:     u.Email.String,
				CreatedAt: u.CreatedAt.Time.Format("2006-01-02T15:04:05-07:00"),
			})
		}
	}

	var res = &model.ResponseWithMultipleData{
		Data:         users,
		Type:         "success",
		Status:       200,
		Message:      "success",
		PageNumber:   data.PageNumber,
		BatchNumber:  data.BatchNumber,
		TotalRecords: totalRecords,
	}

	return res, 200, nil
}

// GetSingleUserController retrieves a single user's data based on the specified user ID.
func GetSingleUserController(id string) (r *model.ResponseWithSingleData, code int, err error) {

	// Create a database connection
	db := Database.Connect()
	defer db.Close()

	// SQL statement to fetch user data along with profile, address, and phone numbers for a specific user ID
	sqlStatement := fmt.Sprintf(`
		SELECT 
			U.id, U.first_name, U.last_name, U.email, U.created_at

		FROM 
			users AS 

		WHERE
			U.id = '%v' 
			AND U.deleted = false;

	`,
		id,
	)

	// Execute the SQL query
	rows, err := db.Query(sqlStatement)
	if err != nil {
		errorLogger.CaptureException("GetSingleUserController--1", fmt.Errorf("%v | error: %v", sqlStatement, err))

		return nil, 500, err
	}
	defer rows.Close()

	// Initialize variables to store user-related data
	user := model.User{}

	// Iterate through query results
	for rows.Next() {
		// Scan data from the database into corresponding structs
		u := model.UserSQL{}

		// Populate user-related structs
		if err := rows.Scan(
			&u.Id, &u.FirstName, &u.LastName, &u.Email, &u.CreatedAt,
		); err != nil {
			errorLogger.CaptureException("GetSingleUserController--2", err)

			return nil, 500, err
		}

		user = model.User{
			Id:        u.Id.String,
			FirstName: u.FirstName.String,
			LastName:  u.LastName.String,
			Email:     u.Email.String,
			CreatedAt: u.CreatedAt.Time.Format("2006-01-02T15:04:05-07:00"),
		}
	}

	// Check if the user ID is empty, indicating that the user was not found
	if user.Id == "" {
		return nil, 400, fmt.Errorf("user %v not found", id)
	}

	// Create a response with the retrieved user data
	var res = &model.ResponseWithSingleData{
		Data:    user,
		Type:    "success",
		Status:  200,
		Message: "success",
	}

	return res, 200, nil
}

// UpdateUserController updates user-related information based on the provided user ID, current user ID, and user account data.
func UpdateUserController(id, currentUser string, data *model.UserFull) (r *model.ResponseWithSingleData, code int, err error) {

	// Create a database connection
	db := Database.Connect()
	defer db.Close()

	// SQL statement to check if the user with the specified ID exists and is not deleted
	sqlStatement := fmt.Sprintf(`
		SELECT 
			id 
			
		FROM 
			users 
			
		WHERE id = '%v' 
			AND deleted = false;
	`,
		id,
	)

	// Execute the SQL query
	rows, err := db.Query(sqlStatement)
	if err != nil {
		errorLogger.CaptureException("UpdateUserController--1", fmt.Errorf("%v | error: %v", sqlStatement, err))

		return nil, 500, err
	}
	defer rows.Close()

	// Check if the user with the specified ID exists
	if !rows.Next() {
		return nil, 400, fmt.Errorf("user %v not found", id)
	}

	// Initialize an empty update statement
	updateStatement := ""

	// Check if there are updates to the user's first name
	if data.FirstName != "" {
		// Build the update statement for the 'users' table
		updateStatement = fmt.Sprintf("%v UPDATE users SET last_modifier_id = '%v', first_name = '%v' WHERE id = '%v';", updateStatement, currentUser, data.FirstName, id)
	}

	// Check if there are updates to the user's last name
	if data.LastName != "" {
		// Build the update statement for the 'users' table
		updateStatement = fmt.Sprintf("%v UPDATE users SET last_modifier_id = '%v', last_name = '%v' WHERE id = '%v';", updateStatement, currentUser, data.LastName, id)
	}

	// Check if there are any updates to be executed
	if updateStatement != "" {

		// Begin a database transaction
		tx, err := db.Begin()
		if err != nil {
			errorLogger.CaptureException("UpdateUserController--2", err)

			return nil, 500, err
		}

		// Execute the update statement within the transaction
		if _, err := tx.Exec(updateStatement); err != nil {

			errorLogger.CaptureException("UpdateUserController--3", fmt.Errorf("%v | error: %v", updateStatement, err))

			// Rollback the transaction if there's an error
			if err := tx.Rollback(); err != nil {
				errorLogger.CaptureException("UpdateUserController--4", err)

				return nil, 500, err
			}

			return nil, 500, err
		}

		// Commit the transaction if there are no errors
		if err := tx.Commit(); err != nil {
			errorLogger.CaptureException("UpdateUserController--5", err)

			return nil, 500, err
		}

		// Create a success response with a message indicating the update
		res := &model.ResponseWithSingleData{
			Data:    struct{}{},
			Type:    "success",
			Status:  200,
			Message: "Updated! NOTE: if you didn't pass id for either user, profile, address, or phone numbers that field will not be updated.",
		}

		// Return an error if there is no data to update
		return res, res.Status, nil
	}

	return nil, 400, errors.New("no data to update")
}

// DeleteUserController deletes a user and associated records based on the provided user ID and current user ID.
func DeleteUserController(id, currentUser string) (r *model.ResponseWithSingleData, code int, err error) {

	// Create a database connection
	db := Database.Connect()
	defer db.Close()

	// SQL statement to check if the user with the specified ID exists and is not deleted
	sqlStatement := fmt.Sprintf(`SELECT id  FROM users WHERE id = '%v' AND deleted = false;`, id)

	// Execute the SQL query
	rows, err := db.Query(sqlStatement)
	if err != nil {
		errorLogger.CaptureException("DeleteUserController--1", fmt.Errorf("%v | error: %v", sqlStatement, err))

		return nil, 500, err
	}
	defer rows.Close()

	// Check if the user with the specified ID exists
	if !rows.Next() {
		return nil, 400, fmt.Errorf("user %v not found", id)
	}

	// Execute multiple SQL statements within a transaction to update/delete user-related records
	updateStatement := fmt.Sprintf(`
		UPDATE 
			users 

		SET 
			deleted = true, 
			last_modifier_id = '%v'

		WHERE 
			id = '%v';

		`,
		currentUser,
		id,
	)

	// Begin a database transaction
	tx, err := db.Begin()
	if err != nil {
		errorLogger.CaptureException("DeleteUserController--2", err)

		return nil, 500, err
	}

	if _, err := tx.Exec(updateStatement); err != nil {

		errorLogger.CaptureException("DeleteUserController--3", fmt.Errorf("%v | error: %v", updateStatement, err))

		// Rollback the transaction if there's an error
		if err := tx.Rollback(); err != nil {
			errorLogger.CaptureException("DeleteUserController--4", err)

			return nil, 500, err
		}

		return nil, 500, err
	}

	// Commit the transaction if there are no errors
	if err := tx.Commit(); err != nil {
		errorLogger.CaptureException("DeleteUserController--5", err)

		return nil, 500, err
	}

	// Create a success response with a message indicating the user deletion
	res := &model.ResponseWithSingleData{
		Data:    struct{}{},
		Type:    "success",
		Status:  200,
		Message: "User Deleted",
	}

	return res, res.Status, nil
}
