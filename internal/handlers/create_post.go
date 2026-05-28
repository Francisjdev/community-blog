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

func RegisterPost(app *app.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		claims, ok := r.Context().Value(middleware.ClaimsKey).(*auth.CustomClaims)
		if !ok {
			log.Printf("this is the claims in handler %v: \n", claims)
			helper.RespondWithError(w, 401, "unauthorized")
			return
		}
		id, err := uuid.Parse(claims.Subject)
		if err != nil {
			helper.RespondWithError(w, 401, "unauthorized")
			return
		}
		decoder := json.NewDecoder(r.Body)
		postPayload := createPostRequest{}
		err = decoder.Decode(&postPayload)
		if err != nil {
			helper.RespondWithError(w, 500, "Error decoding parameters")
			log.Printf("Error decoding parameters: %s", err)
			return
		}
		meta := helper.ToNullString(postPayload.MetaDescription)
		cover_image := helper.ToNullString(postPayload.CoverImageUrl)
		youtube_links := helper.ToRawMessage(postPayload.YoutubeLinks)
		published_at := helper.StringToNullTime(postPayload.PublishedAt)

		registerPostParams := database.CreatePostParams{
			Title:           postPayload.Title,
			Slug:            postPayload.Slug,
			MarkdownContent: postPayload.MarkdownContent,
			MetaDescription: meta,
			CoverImageUrl:   cover_image,
			YoutubeLinks:    youtube_links,
			PublishedAt:     published_at,
			UserID:          id,
		}
		returnedPostData, err := app.Service.Posts.RegisterPost(r.Context(), registerPostParams)
		if err != nil {
			helper.RespondWithError(w, 500, "Error creating post")
			log.Printf("Error creating user: %s", err)
			return
		}
		post := postResponse{
			Title:       returnedPostData.Title,
			Slug:        returnedPostData.Slug,
			PublishedAt: returnedPostData.PublishedAt,
		}
		helper.RespondWithJSON(w, 201, post)
	}
}
