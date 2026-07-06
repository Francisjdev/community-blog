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
-- name: DeletePostById :exec
DELETE FROM posts
WHERE id =$1 and user_id = $2;

-- name: DeleteAllPostByUser :exec
DELETE FROM posts
WHERE user_id =$1;

-- name: GetPostById :one
SELECT * FROM posts
WHERE id = $1;


-- name: ListAllPostsByUser :many
SELECT * FROM posts
WHERE user_id = $1;

-- name: ListAllPosts :many
SELECT * FROM posts;
