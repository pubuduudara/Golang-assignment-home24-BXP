package middleware

import (
	"net/http"
	"os"

	"github.com/pubuduudara/Golang-assignment-home24-BXP/backend/internal/utils/logger"
)

// APIKeyMiddleware to ensures valid API key is sent in the request header
func APIKeyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		apiKey := req.Header.Get("X-API-Key")
		expectedKey := os.Getenv("API_KEY")

		if apiKey == "" || apiKey != expectedKey {
			res.WriteHeader(http.StatusUnauthorized)
			res.Write([]byte(`{"status": false, "data": "Unauthorized - Invalid API key"}`))
			logger.Warn("Unauthorized access attempt")
			return
		}
		next.ServeHTTP(res, req)
	})
}
