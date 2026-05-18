package main

import (
	"github.com/francisjdev/community-blog/internal/app"
	"github.com/francisjdev/community-blog/internal/handlers"
	"github.com/go-chi/chi/v5"
)

func LoadRouter(app *app.Application) *chi.Mux {
	r := chi.NewRouter()
	r.Post("/api/users", handlers.RegisterUser(app))
	return r
}
