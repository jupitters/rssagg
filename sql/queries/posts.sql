-- name: CreatePost :one
INSERT INTO users(id, created_at, updated_at, title, description, published_at, url, feed_id)
VALUES($1, $2, $3, $4, $5, $6, %7, %8)
RETURNING *;

-- name: GetUserByApiKey :one
SELECT * FROM users WHERE api_key = $1;