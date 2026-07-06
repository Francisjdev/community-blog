package handlers

import (
	"log"
	"net/http"

	"github.com/francisjdev/community-blog/internal/app"
	"github.com/francisjdev/community-blog/internal/helper"
)

func GetPosts(app *app.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		posts, err := app.Service.Posts.ListAllPosts(r.Context())
		if err != nil {
			helper.RespondWithError(w, 500, "Error getting all post")
			log.Printf("Error getting all posts: %s", err)
			return
		}

		helper.RespondWithJSON(w, 200, posts)

	}
}
