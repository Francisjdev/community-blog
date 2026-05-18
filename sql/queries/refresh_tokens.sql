-- name: CreateRefreshToken :one
INSERT INTO refresh_tokens (
  user_id,
  token_hash
) VALUES (
  $1, $2
)
RETURNING *;
