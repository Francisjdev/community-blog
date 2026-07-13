package handlers

import (
	"log"
	"net/http"

	"github.com/francisjdev/community-blog/internal/app"
	"github.com/francisjdev/community-blog/internal/auth"
	"github.com/francisjdev/community-blog/internal/database"
	"github.com/francisjdev/community-blog/internal/helper"
	"github.com/francisjdev/community-blog/internal/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func DeleteSinglePost(app *app.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		claims, ok := r.Context().Value(middleware.ClaimsKey).(*auth.CustomClaims)
		if !ok {
			helper.RespondWithError(w, 401, "unauthorized")
			return
		}
		userID, err := uuid.Parse(claims.Subject)
		if err != nil {
			helper.RespondWithError(w, 400, "Bad Request")
			return
		}
		id := chi.URLParam(r, "id")
		postId, err := uuid.Parse(id)
		deleteParams := database.DeletePostByIdParams{
			ID:     postId,
			UserID: userID,
		}
		err = app.Service.Posts.DeleteSinglePost(r.Context(), deleteParams)
		if err != nil {
			helper.RespondWithError(w, 500, "Error Deleting post")
			log.Printf("Error creating user: %s", err)
			return
		}
		helper.RespondWithJSON(w, 200, map[string]string{"message": "Post deleted successfully"})
	}
}
func DeleteAllPostByUser(app *app.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		claims, ok := r.Context().Value(middleware.ClaimsKey).(*auth.CustomClaims)
		if !ok {
			helper.RespondWithError(w, 401, "unauthorized")
			return
		}
		//finish this soon
		_, err := uuid.Parse(claims.Subject)
		if err != nil {
			helper.RespondWithError(w, 400, "Bad Request")
			return
		}
	}
}
