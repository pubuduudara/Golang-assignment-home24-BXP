package api

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

// integration test for the analyze endpoint success workflow
func TestIntegration_Analyze_Success(t *testing.T) {
	os.Setenv("API_KEY", "test-key")

	router := MainRouter()

	// Build request
	req := httptest.NewRequest(http.MethodGet, "/analyze?url=https://httpstat.us/200", nil)
	req.Header.Set("X-API-Key", "test-key")

	// Record response
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected 200 OK, got %d", rr.Code)
	}

	if !strings.Contains(rr.Body.String(), `"status":true`) {
		t.Errorf("Expected success response body, got: %s", rr.Body.String())
	}
}

// Missing URL
func TestIntegration_Analyze_MissingURL(t *testing.T) {
	os.Setenv("API_KEY", "test-key")

	router := MainRouter()
	req := httptest.NewRequest(http.MethodGet, "/analyze", nil)
	req.Header.Set("X-API-Key", "test-key")

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("Expected 400 Bad Request, got %d", rr.Code)
	}

	if !strings.Contains(rr.Body.String(), "Missing 'url'") {
		t.Errorf("Expected error message, got: %s", rr.Body.String())
	}
}

// No API key
func TestIntegration_Analyze_Unauthorized(t *testing.T) {
	router := MainRouter()
	req := httptest.NewRequest(http.MethodGet, "/analyze?url=https://example.com", nil)

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusUnauthorized {
		t.Errorf("Expected 401 Unauthorized, got %d", rr.Code)
	}

	if !strings.Contains(rr.Body.String(), "Unauthorized") {
		t.Errorf("Expected unauthorized message, got: %s", rr.Body.String())
	}
}
