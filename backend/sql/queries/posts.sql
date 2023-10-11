-- name: CreatePost :one
INSERT INTO posts (id, created_at, updated_at, title, url, description, published_at, feed_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: GetPostsByUser :many
SELECT * FROM
posts P JOIN feeds F ON P.feed_id = F.feed_id
WHERE user_id = $1
ORDER BY P.created_at
LIMIT $2;