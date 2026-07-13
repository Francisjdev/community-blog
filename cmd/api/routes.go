package main

import (
	"github.com/francisjdev/community-blog/internal/app"
	"github.com/francisjdev/community-blog/internal/handlers"
	"github.com/francisjdev/community-blog/internal/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func LoadRouter(app *app.Application) *chi.Mux {
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	r.Post("/api/users", handlers.RegisterUser(app))
	r.Delete("/api/users", middleware.AuthMiddleware(app, handlers.DeleteUser(app)))
	r.Post("/api/auth/login", handlers.Login(app))
	r.Post("/api/posts/create", middleware.AuthMiddleware(app, handlers.RegisterPost(app)))
	r.Delete("/api/posts/{id}", middleware.AuthMiddleware(app, handlers.DeleteSinglePost(app)))
	r.Get("/api/posts", handlers.GetPosts(app))
	r.Get("/api/posts/{id}", handlers.GetPostByID(app))

	return r
}
