package main

import (
	"net/http"

	"github.com/joho/godotenv"
	"github.com/pubuduudara/Golang-assignment-home24-BXP/backend/internal/api"
	"github.com/pubuduudara/Golang-assignment-home24-BXP/backend/internal/utils/logger"
)

func main() {
	if err := godotenv.Load(); err != nil {
		logger.Error(err, "Error loading .env file, please check if the file exists. If not, create a new .env file with the required environment variables.")
		return
	}
	port := ":8080"
	logger.Info("starting web server on port " + port)

	router := api.MainRouter()

	err := http.ListenAndServe(port, router)
	if err != nil {
		logger.Error(err, "failed to start server")
	}
}
