package main

import (
	"log"
	"os"

	view "ai-project/views"

	"github.com/joho/godotenv"
)

func main() {

	ENV := os.Getenv("APP_ENV")

	if ENV != "staging" || ENV != "qa" && ENV != "production" {
		godotenv.Load(".env")
	}

	DB_HOST := os.Getenv("DB_HOST")

	if DB_HOST == "" {
		if ENV == "staging" || ENV == "qa" || ENV == "production" {
			os.Exit(1)
		} else {
			log.Fatal("Unable to load DB_HOST")
		}
	}

	DB_USER := os.Getenv("DB_USER")

	if DB_USER == "" {
		if ENV == "staging" || ENV == "qa" || ENV == "production" {
			os.Exit(1)
		} else {
			log.Fatal("Unable to load DB_USER")
		}
	}

	DB_PASSWORD := os.Getenv("DB_PASSWORD")

	if DB_PASSWORD == "" {
		if ENV == "staging" || ENV == "qa" || ENV == "production" {
			os.Exit(1)
		} else {
			log.Fatal("Unable to load DB_PASSWORD")
		}
	}

	DB_NAME := os.Getenv("DB_NAME")

	if DB_NAME == "" {
		if ENV == "staging" || ENV == "qa" || ENV == "production" {
			os.Exit(1)
		} else {
			log.Fatal("Unable to load DB_NAME")
		}
	}

	DB_PORT := os.Getenv("DB_PORT")

	if DB_PORT == "" {
		if ENV == "staging" || ENV == "qa" || ENV == "production" {
			os.Exit(1)
		} else {
			log.Fatal("Unable to load DB_PORT")
		}
	}

	APP_PORT := os.Getenv("APP_PORT")
	if APP_PORT == "" {
		if ENV == "staging" || ENV == "qa" || ENV == "production" {
			os.Exit(1)
		} else {
			log.Fatal("Unable to load APP_PORT")
		}
	}

	view.StartServer()

}
