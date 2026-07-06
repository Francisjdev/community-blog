package main

import (
	"github.com/francisjdev/community-blog/internal/app"
	"github.com/francisjdev/community-blog/internal/handlers"
	"github.com/francisjdev/community-blog/internal/middleware"
	"github.com/go-chi/chi/v5"
)

func LoadRouter(app *app.Application) *chi.Mux {
	r := chi.NewRouter()
	r.Post("/api/users", handlers.RegisterUser(app))
	r.Delete("/api/users", middleware.AuthMiddleware(app, handlers.DeleteUser(app)))
	r.Post("/api/auth/login", handlers.Login(app))
	r.Post("/api/posts/create", middleware.AuthMiddleware(app, handlers.RegisterPost(app)))
	r.Post("/api/posts/deletesinglepost", middleware.AuthMiddleware(app, handlers.DeleteSinglePost(app)))
	r.Get("/api/posts", handlers.GetPosts(app))
	return r
}
