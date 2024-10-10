# AI Car Masters Project

This project focuses on developing and deploying a comprehensive predictive analytics framework for the automotive industry, utilizing machine learning and deep learning techniques. The primary objective is to create an accurate model for predicting car prices based on historical data.

The project culminates in the deployment of the best-performing models through a fully functional web application. This application features an intuitive user interface, a robust backend server, and a database to enable users to input a dataset with the specified vehicle attributes and receive real-time price predictions for viewing and download.

## Project Dependencies

- [Python 3.x.x](https://www.python.org/downloads/)
- [Jupyter Noteboks](https://jupyter.org/install)
- [Pandas](https://pandas.pydata.org/docs/getting_started/install.html)
- [Plotly](https://plotly.com/python/getting-started/)
- [Scikit Learn](https://scikit-learn.org/)
- [Joblib](https://joblib.readthedocs.io/)
- [Go](https://go.dev/doc/install)
- [PostgreSQL](https://www.postgresql.org/download/)
- [Goose](https://github.com/pressly/goose)

## Installation Commands

Assuming you've python 3, Go (Golang) and PostgreSQL installed run the following commands on your terminal/shell/command prompt


pip install jupyterlab



pip install pandas




pip install scikit-learn



pip install plotly



pip install joblib



pip install numpy



go install github.com/pressly/goose/v3/cmd/goose@latest


## For Fresh Start

Make sure you're in the project directory run dir for windows or pwd on mac and linux to confirm


go install github.com/pressly/goose/v3/cmd/goose@latest



cd migration



goose postgres postgres://DB_USER:DB_PASSWORD:@DB_HOST:DB_PORT/DB_NAME up
goose postgres postgres://postgres:postgres@localhost:5432/ai-project up


in your postgreSQL terminal/cmd/shell run CREATE DATABASE ai-project


cd ..


Change the following credentials in the .env file

- DB_HOST
- DB_USER
- DB_PASSWORD


DB_HOST=DB_HOST
DB_PORT=5432
DB_USER=DB_USER
DB_PASSWORD=DB_PASSWORD
DB_NAME=ai-project
DB_URL=postgres://DB_USER:DB_PASSWORD:@DB_HOST:5432/ai-project


## Preset Data

Run this sql file [ai-project-database]("./ai-project-database.sql")

## Dependency Installation

Make sure you're in the project directory run dir for windows or pwd on mac and linux to confirm


go mod tidy



go mod vendor


## Start Command

Make sure you're in the project directory run dir for windows or pwd on mac and linux to confirm


go run main.go


## Build Command

Make sure you're in the project directory run dir for windows or pwd on mac and linux to confirm


go build -o app .


## Run Build Command

Make sure you're in the project directory run dir for windows or pwd on mac and linux to confirm

On Mac and Linux

./app

On Windows

.\app