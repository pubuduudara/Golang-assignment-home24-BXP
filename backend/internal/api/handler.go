package api

import (
	"encoding/json"
	"net/http"

	"github.com/pubuduudara/Golang-assignment-home24-BXP/backend/internal/models"
	"github.com/pubuduudara/Golang-assignment-home24-BXP/backend/internal/services"
	"github.com/pubuduudara/Golang-assignment-home24-BXP/backend/internal/utils/error"
	"github.com/pubuduudara/Golang-assignment-home24-BXP/backend/internal/utils/helpers"
	"github.com/pubuduudara/Golang-assignment-home24-BXP/backend/internal/utils/logger"
)

// handles the /analyze route
func analyzeHandler(res http.ResponseWriter, req *http.Request) {
	queryURL := req.URL.Query().Get("url")
	if queryURL == "" {
		writeError(res, "Missing 'url' query parameter", http.StatusBadRequest)
		return
	}

	if !helpers.IsValidURL(queryURL) {
		writeError(res, "Invalid URL format", http.StatusBadRequest)
		return
	}

	result, err := services.AnalyzeURL(queryURL)
	if err != nil {
		if e, ok := err.(*services.RequestError); ok {
			error.RespondWithError(res, e.StatusCode, e.Error())
			return
		}
		// for other unexpected errors, return internal server error
		error.RespondWithError(res, http.StatusInternalServerError, "Internal Server Error: "+err.Error())
		return
	}
	response := models.Response{
		Status: true,
		Data:   result,
	}

	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(response)
}

func writeError(w http.ResponseWriter, message string, code int) {
	logger.Error(nil, message)
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": false,
		"data":   message,
	})
}
