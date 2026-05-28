package handlers

import (
	"database/sql"
)

type User struct {
	Email string `json:"email"`
}

type createUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type createPostRequest struct {
	Title           string `json:"title"`
	Slug            string `json:"slug"`
	MarkdownContent string `json:"markdown_content"`
	MetaDescription string `json:"meta_description"`
	CoverImageUrl   string `json:"cover_image_url"`
	YoutubeLinks    string `json:"youtube_links"`
	PublishedAt     string `json:"published_at"`
	UserID          string `json:"user_id"`
}

type postResponse struct {
	Title       string       `json:"title"`
	Slug        string       `json:"slug"`
	PublishedAt sql.NullTime `json:"published_at"`
}
