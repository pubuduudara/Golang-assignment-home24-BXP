package api

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func MainRouter() http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	return router

}
