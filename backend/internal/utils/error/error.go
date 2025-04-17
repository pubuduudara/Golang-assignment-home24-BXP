package error

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pubuduudara/Golang-assignment-home24-BXP/backend/internal/utils/logger"
)

// writes a structured JSON error response
func RespondWithError(w http.ResponseWriter, status int, msg string) {
	logger.Error(nil, msg)
	formattedMsg := fmt.Sprintf("Failed with error code %d: %s", status, msg)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": false,
		"data":   formattedMsg,
	})
}
