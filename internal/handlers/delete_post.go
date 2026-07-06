package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/francisjdev/community-blog/internal/app"
	"github.com/francisjdev/community-blog/internal/auth"
	"github.com/francisjdev/community-blog/internal/database"
	"github.com/francisjdev/community-blog/internal/helper"
	"github.com/francisjdev/community-blog/internal/middleware"
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
		decoder := json.NewDecoder(r.Body)
		Payload := deletePostRequest{}
		err = decoder.Decode(&Payload)
		if err != nil {
			helper.RespondWithError(w, 500, "Error decoding parameters")
			log.Printf("Error decoding parameters: %s", err)
			return
		}
		postID, err := uuid.Parse(Payload.PostId)
		if err != nil {
			helper.RespondWithError(w, 400, "Bad Request")
			return
		}
		deleteParams := database.DeletePostByIdParams{
			ID:     postID,
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
