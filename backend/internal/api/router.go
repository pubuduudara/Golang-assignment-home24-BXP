package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/pubuduudara/Golang-assignment-home24-BXP/backend/internal/middleware"
)

func MainRouter() http.Handler {
	router := chi.NewRouter()
	// add CORS middleware first
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"}, // React dev server
		AllowedMethods: []string{"GET"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-API-Key"},
	}))
	// add API key middleware
	router.Use(middleware.APIKeyMiddleware)
	router.Get("/analyze", analyzeHandler)
	return router
}
