package main

import "github.com/go-chi/chi"

func (a *Application) routes() *chi.Mux {
	r := chi.NewRouter()

	// add cors

	r.Post("/authenticate", a.Auth)
	return r
}
