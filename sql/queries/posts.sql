-- name: CreatePost :one
INSERT INTO posts (
  title,
  slug,
  markdown_content,
  meta_description,
  cover_image_url,
  youtube_links,
  published_at,
  user_id
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8
)
RETURNING *;
