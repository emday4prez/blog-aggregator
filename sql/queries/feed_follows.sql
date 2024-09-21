-- name: CreateFeedFollow :one
INSERT INTO feed_follows (id,created_at,updated_at,user_id ,feed_id)
VALUES ($1, $2, $3, $4, encode(sha256(random()::text::bytea), 'hex'))
RETURNING id, created_at, updated_at, user_id, feed_id;