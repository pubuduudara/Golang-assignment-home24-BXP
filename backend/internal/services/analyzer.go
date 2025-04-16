package services

import (
	"net/http"

	"github.com/pubuduudara/Golang-assignment-home24-BXP/backend/internal/models"
)

// RequestError is used to handle errors from HTTP requests
// TODO: define this in models
type RequestError struct {
	StatusCode int
}

// service function for analyzing URLs
func AnalyzeURL(targetURL string) (*models.PageAnalysis, error) {
	return &models.PageAnalysis{
		HTMLVersion:  "",
		Title:        "",
		Headings:     map[string]int{},    // to be filled in later
		Links:        models.LinkCounts{}, // to be filled in later
		HasLoginForm: false,               // to be filled in later
	}, nil
}

func (e *RequestError) Error() string {
	return http.StatusText(e.StatusCode)
}
