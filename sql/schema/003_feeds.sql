-- +goose Up
CREATE TABLE feeds(
 id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
 url TEXT UNIQUE,
 user_id UUID REFERENCES users(id) ON DELETE CASCADE,
);
-- +goose Down
DROP TABLE feeds;