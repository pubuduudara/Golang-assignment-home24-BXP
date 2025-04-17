package middleware

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/pubuduudara/Golang-assignment-home24-BXP/backend/internal/utils/logger"
)

// test API key middleware
func TestAPIKeyMiddleware(t *testing.T) {
	const validKey = "test-api-key"
	os.Setenv("API_KEY", validKey)

	tests := []struct {
		name       string
		apiKey     string
		wantStatus int
		wantBody   string
	}{
		{
			name:       "valid API key",
			apiKey:     validKey,
			wantStatus: http.StatusOK,
			wantBody:   "authorized",
		},
		{
			name:       "missing API key",
			apiKey:     "",
			wantStatus: http.StatusUnauthorized,
			wantBody:   `{"status": false, "data": "Unauthorized - Invalid API key"}`,
		},
		{
			name:       "invalid API key",
			apiKey:     "wrong-key",
			wantStatus: http.StatusUnauthorized,
			wantBody:   `{"status": false, "data": "Unauthorized - Invalid API key"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a test handler to wrap with middleware
			finalHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("authorized"))
			})

			// Wrap the handler with middleware
			handler := APIKeyMiddleware(finalHandler)

			// Build request
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			if tt.apiKey != "" {
				req.Header.Set("X-API-Key", tt.apiKey)
			}

			// Capture response
			rec := httptest.NewRecorder()
			handler.ServeHTTP(rec, req)

			// Assert status code
			if rec.Code != tt.wantStatus {
				logger.Error(nil, "Wrong status code")
				t.Errorf("got status %d, want %d", rec.Code, tt.wantStatus)
			}

			// Assert body
			body := strings.TrimSpace(rec.Body.String())
			if body != tt.wantBody {
				logger.Error(nil, "Wrong response body")
				t.Errorf("got body %q, want %q", body, tt.wantBody)
			}
		})
	}
}
