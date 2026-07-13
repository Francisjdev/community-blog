package handlers

import (
	"log"
	"net/http"

	"github.com/francisjdev/community-blog/internal/app"
	"github.com/francisjdev/community-blog/internal/helper"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func GetPosts(app *app.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		posts, err := app.Service.Posts.ListAllPosts(r.Context())
		if err != nil {
			helper.RespondWithError(w, 500, "Error getting all post")
			log.Printf("Error getting all posts: %s", err)
			return
		}
		formattedPosts := []postResponse{}
		for _, val := range posts {
			publishedAt := ""
			if val.PublishedAt.Valid {
				publishedAt = val.PublishedAt.Time.String()
			}
			youtubeLinks := ""
			if val.YoutubeLinks.Valid {
				youtubeLinks = string(val.YoutubeLinks.RawMessage)
			}
			postr := postResponse{
				Title:           val.Title,
				Slug:            val.Slug,
				MarkdownContent: val.MarkdownContent,
				MetaDescription: val.MetaDescription.String,
				CoverImageUrl:   val.CoverImageUrl.String,
				YoutubeLinks:    youtubeLinks,
				PublishedAt:     publishedAt,
				UserID:          val.UserID.String(),
				PostID:          val.ID.String(),
			}
			formattedPosts = append(formattedPosts, postr)
		}
		helper.RespondWithJSON(w, 200, formattedPosts)

	}
}

func GetPostByID(app *app.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		//need to make sure how to handle error when id is not existent
		postId, err := uuid.Parse(id)
		if err != nil {
			helper.RespondWithError(w, 400, "bad request")
			return
		}

		post, err := app.Service.Posts.GetPostById(r.Context(), postId)
		if err != nil {
			helper.RespondWithError(w, 400, "bad request")
			return
		}
		publishedAt := ""
		if post.PublishedAt.Valid {
			publishedAt = post.PublishedAt.Time.String()
		}
		youtubeLinks := ""
		if post.YoutubeLinks.Valid {
			youtubeLinks = string(post.YoutubeLinks.RawMessage)
		}

		responsePost := postResponse{
			Title:           post.Title,
			Slug:            post.Slug,
			MarkdownContent: post.MarkdownContent,
			MetaDescription: post.MetaDescription.String,
			CoverImageUrl:   post.CoverImageUrl.String,
			YoutubeLinks:    youtubeLinks,
			PublishedAt:     publishedAt,
			UserID:          post.UserID.String(),
			PostID:          post.ID.String(),
		}
		helper.RespondWithJSON(w, 200, responsePost)
	}
}
