package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

func (a *Application) routes() *chi.Mux {
	r := chi.NewRouter()

	// TODO cors
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://*", "https://*"},
		AllowedMethods: []string{"POST", "GET", "PUT", "DELETE", "OPTION"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	}))

	r.Use(middleware.Heartbeat("/ping"))
	r.Get("/handle", a.handleSubmission)
	return r
}
