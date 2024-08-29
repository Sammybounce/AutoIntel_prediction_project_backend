package Database

import (
	errorLogger "ai-project/utils/error-logger"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

type Database struct {
	DB *sql.DB
}

func (database *Database) init(host, user, password, dbName string, port int) {

	// connection string
	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)

	// open database
	db, err := sql.Open("postgres", connString)
	checkError(err)

	// check db
	err = db.Ping()
	checkError(err)

	database.DB = db
}

func checkError(err error) {

	if err != nil {
		errorLogger.CaptureException("checkError", err)
		os.Exit(1)
	}
}

func Connect() *sql.DB {
	database := &Database{}

	DB_HOST := os.Getenv("DB_HOST")
	DB_USER := os.Getenv("DB_USER")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_NAME := os.Getenv("DB_NAME")
	DB_PORT := os.Getenv("DB_PORT")

	PORT, err := strconv.Atoi(DB_PORT)

	if err != nil {
		log.Fatal("Unable to DB_PORT to int")
	}

	database.init(DB_HOST, DB_USER, DB_PASSWORD, DB_NAME, PORT)

	return database.DB
}
