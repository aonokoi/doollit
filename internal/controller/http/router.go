package http

import (
	"net/http"

	"proj/doollit/internal/usecase"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Router(service *usecase.STask) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.URLFormat)
	r.Use(middleware.RequestID)

	handlers := NewHandlers(service)

	r.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Post("/task", handlers.CreateTask)
		})
	})

	return r
}
