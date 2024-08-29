package controller

import (
	"database/sql"
	"fmt"

	model "ai-project/models"
	customStrings "ai-project/utils/custom-strings"
	customTime "ai-project/utils/custom-time"
	Database "ai-project/utils/database"
	errorLogger "ai-project/utils/error-logger"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// SignInController handles user sign-in based on the provided sign-in data.
func SignInController(data *model.SignIn) (r *model.ResponseWithSingleData, code int, err error) {

	// Create a database connection
	db := Database.Connect()
	defer db.Close()

	// SQL statement to retrieve user ID and hashed password for the given email (if not deleted)
	sqlStatement := fmt.Sprintf(`	
		SELECT 
			id, 
			first_name, 
			last_name, 
			u_password

		FROM 
			users

		WHERE 
			email = LOWER('%v') 
			AND deleted = false;
	`,
		data.Email,
	)

	// Execute the SQL query
	rows, err := db.Query(sqlStatement)
	if err != nil {
		errorLogger.CaptureException("SignInController--1", fmt.Errorf("%v | error: %v", sqlStatement, err))
		return nil, 500, err
	}
	defer rows.Close()

	var userId, firstName, lastName, password sql.NullString

	// Retrieve user ID and hashed password from the query result
	for rows.Next() {
		if err := rows.Scan(&userId, &firstName, &lastName, &password); err != nil {
			errorLogger.CaptureException("SignInController--2", err)
			return nil, 500, err
		}
	}

	// Check if the user ID or password is empty, or if the provided password is invalid
	if userId.String == "" {
		return nil, 400, fmt.Errorf("invalid Email or password")
	}

	if userId.String == "" || !checkPasswordHash(data.Password, password.String) {
		return nil, 400, fmt.Errorf("invalid Email or password")
	}

	// Generate a new UUID for the user token
	id, err := uuid.NewV7()
	if err != nil {
		return nil, 500, err
	}

	// Create a new user token with a generated token, expiration time, and other details
	userToken := model.UserToken{
		Id:       fmt.Sprintf("%v", id),
		UserId:   userId.String,
		Token:    customStrings.GenerateRandomString(50),
		ExpireAt: customTime.AddMinutesToCurrentTime(15).Format("2006-01-02T15:04:05-07:00"),
	}

	// SQL statement to insert the user token into the 'user_tokens' table
	insertStatement := fmt.Sprintf(`
		INSERT INTO user_tokens
			(id, user_id, token, expire_at)
		VALUES
			('%v', '%v', '%v', '%v')
	
		RETURNING 
			expire_at;
	`,
		userToken.Id, userToken.UserId, userToken.Token, userToken.ExpireAt,
	)

	rows, err = db.Query(insertStatement)
	if err != nil {
		errorLogger.CaptureException("SignInController--4", fmt.Errorf("%v | error: %v", insertStatement, err))

		return nil, 500, err
	}
	// Retrieve details of the inserted user token from the query result
	for rows.Next() {
		if err := rows.Scan(&userToken.ExpireAt); err != nil {
			errorLogger.CaptureException("SignInController--5", err)
			return nil, 500, err
		}
	}

	// Create a success response with the generated user token
	res := &model.ResponseWithSingleData{
		Data:    userToken,
		Type:    "success",
		Status:  200,
		Message: "success",
	}

	return res, res.Status, nil
}

// SignUpController handles user sign-up based on the provided sign-up data.
func SignUpController(data *model.SignUp) (r *model.ResponseWithSingleData, code int, err error) {

	// Create a database connection
	db := Database.Connect()
	defer db.Close()

	// Check if the provided email already exists and is not deleted
	queryStatement := fmt.Sprintf(`
		SELECT 
			email 
			
		FROM 
			users 
			
		WHERE 
			email = LOWER('%v') 
			AND deleted = false;
	`,
		data.Email,
	)

	rows, err := db.Query(queryStatement)
	if err != nil {
		errorLogger.CaptureException("SignUpController--1", fmt.Errorf("%v | error: %v", queryStatement, err))

		return nil, 500, err
	}
	defer rows.Close()

	// If the email already exists, return an error
	if rows.Next() {
		return nil, 400, fmt.Errorf("email %v already exist", data.Email)
	}

	// Hash the user's password for storage in the database
	hashPassword, err := hashPassword(data.Password)
	if err != nil {
		errorLogger.CaptureException("SignUpController--2", err)

		return nil, 500, err
	}

	userId, err := uuid.NewV7()
	if err != nil {
		return nil, 500, err
	}

	// SQL statement to insert user and user profile data into the database
	sqlStatement := fmt.Sprintf(`
		INSERT INTO users
			(id, first_name, last_name, email, u_password)
		VALUES
			('%v', LOWER('%v'), LOWER('%v'), LOWER('%v'), '%v');
	`,
		userId, data.FirstName, data.LastName, data.Email, hashPassword,
	)

	// Begin a database transaction
	tx, err := db.Begin()
	if err != nil {
		errorLogger.CaptureException("SignUpController--3", err)

		return nil, 500, err
	}

	// Execute the SQL statement within the transaction
	if _, err := tx.Exec(sqlStatement); err != nil {
		errorLogger.CaptureException("SignUpController--6", fmt.Errorf("%v | error: %v", sqlStatement, err))

		// Rollback the transaction if there's an error
		if err := tx.Rollback(); err != nil {
			errorLogger.CaptureException("SignUpController--7", err)

			return nil, 500, err
		}

		return nil, 500, err
	}

	// Commit the transaction if there are no errors
	if err := tx.Commit(); err != nil {
		errorLogger.CaptureException("SignUpController--8", err)

		return nil, 500, err
	}

	// Create a success response with a message and the generated verification code
	res := &model.ResponseWithSingleData{
		Data:    struct{}{},
		Type:    "success",
		Status:  200,
		Message: "success",
	}

	return res, res.Status, nil
}

// AuthenticateController authenticates a user based on the provided token.
func AuthenticateController(token string) (r *model.ResponseWithSingleData, code int, err error) {

	// Create a database connection
	db := Database.Connect()
	defer db.Close()

	// SQL statement to retrieve user token information for the given token
	sqlStatement := fmt.Sprintf(`	
		SELECT 
			id, 
			user_id, 
			token, 
			expire_at 
			
		FROM 
			user_tokens

		WHERE 
			token = '%v' 
			AND expire_at > CURRENT_TIMESTAMP;
	`, token)

	// Execute the SQL query to retrieve user token information
	rows, err := db.Query(sqlStatement)
	if err != nil {
		errorLogger.CaptureException("AuthenticateController--1", fmt.Errorf("%v | error: %v", sqlStatement, err))

		return nil, 500, err
	}
	defer rows.Close()

	// Initialize a UserToken struct to store the retrieved token information
	userToken := model.UserToken{}

	// Extract token information from the query result
	for rows.Next() {
		if err := rows.Scan(&userToken.Id, &userToken.UserId, &userToken.Token, &userToken.ExpireAt); err != nil {
			errorLogger.CaptureException("AuthenticateController--2", err)

			return nil, 500, err
		}
	}

	// If no token information is found, return an error indicating that the token does not exist or has expired
	if userToken.Id == "" {
		return nil, 401, fmt.Errorf("token Does not exist or has expired")
	}

	// Create a success response with the retrieved user token information
	res := &model.ResponseWithSingleData{
		Data:    userToken,
		Type:    "success",
		Status:  200,
		Message: "success",
	}

	return res, res.Status, nil
}

// RefreshTokenController refreshes a user token based on the provided token.
func RefreshTokenController(token string) (r *model.ResponseWithSingleData, code int, err error) {

	// Create a database connection
	db := Database.Connect()
	defer db.Close()

	// SQL statement to retrieve user token information for the given token
	sqlStatement := fmt.Sprintf(`	
		SELECT 
			id, 
			user_id, 
			token, 
			expire_at 
			
		FROM 
			user_tokens

		WHERE 
			token = '%v' 
			AND expire_at > CURRENT_TIMESTAMP;
	`,
		token,
	)

	// Execute the SQL query to retrieve user token information
	rows, err := db.Query(sqlStatement)
	if err != nil {
		errorLogger.CaptureException("RefreshTokenController--1", fmt.Errorf("%v | error: %v", sqlStatement, err))

		return nil, 500, err
	}
	defer rows.Close()

	// Initialize a UserToken struct to store the retrieved token information
	userToken := model.UserToken{}

	// Extract token information from the query result
	for rows.Next() {
		if err := rows.Scan(&userToken.Id, &userToken.UserId, &userToken.Token, &userToken.ExpireAt); err != nil {
			errorLogger.CaptureException("RefreshTokenController--2", err)

			return nil, 500, err
		}
	}

	// If no token information is found, return an error indicating that the token is invalid
	if userToken.Id == "" {
		return nil, 401, fmt.Errorf("token does not exist or has expired")
	}

	// Execute an update statement to refresh the token expiration time
	execStatement := fmt.Sprintf(`
		UPDATE 
			user_tokens

		SET 
			expire_at = '%v'

		WHERE 
			token = '%v';

		------------

		SELECT 
			id, 
			user_id, 
			token, 
			expire_at 
			
		FROM 
			user_tokens

		WHERE 
			token = '%v';
		`,
		customTime.AddMinutesToCurrentTime(15).Format("2006-01-02T15:04:05-07:00"),
		token,
		token,
	)

	// Begin a database transaction to update the token expiration time
	tx, err := db.Begin()
	if err != nil {
		errorLogger.CaptureException("RefreshTokenController--3", err)

		return nil, 500, err
	}

	rows, err = tx.Query(execStatement)
	if err != nil {
		errorLogger.CaptureException("RefreshTokenController--4", err)

		// Rollback the transaction if there's an error
		if err := tx.Rollback(); err != nil {
			errorLogger.CaptureException("RefreshTokenController--5", err)

			return nil, 500, err
		}

		return nil, 500, err
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		errorLogger.CaptureException("RefreshTokenController--6", err)

		return nil, 500, err
	}

	var expireAt sql.NullTime

	// Extract the refreshed token information from the transaction result
	for rows.Next() {
		if err := rows.Scan(&userToken.Id, &userToken.UserId, &userToken.Token, &expireAt); err != nil {
			errorLogger.CaptureException("RefreshTokenController--7", err)

			return nil, 500, err
		}
	}

	userToken.ExpireAt = expireAt.Time.Format("2006-01-02T15:04:05-07:00")

	// Create a success response with the refreshed user token information
	res := &model.ResponseWithSingleData{
		Data:    userToken,
		Type:    "success",
		Status:  200,
		Message: "success",
	}

	return res, res.Status, nil
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		errorLogger.CaptureException("hashPassword--1", err)
		return "", err
	}

	return string(hashedPassword), nil
}

func checkPasswordHash(text, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(text))
	if err != nil {
		errorLogger.CaptureException("hashPassword--2", err)
	}

	return err == nil
}
