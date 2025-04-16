package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/pubuduudara/Golang-assignment-home24-BXP/backend/internal/middleware"
)

func MainRouter() http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.APIKeyMiddleware)
	router.Get("/analyze", analyzeHandler)
	return router
}
