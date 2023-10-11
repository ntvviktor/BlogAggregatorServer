-- +goose Up
CREATE table feed_follows(
    id UUID PRIMARY KEY,
    created_at  TIMESTAMP NOT NULL,
    updated_at  TIMESTAMP NOT NULL,
    feed_id UUID NOT NULL REFERENCES feeds(id) ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE, 
    UNIQUE(user_id, feed_id)
);

-- +goose Down
DROP TABLE feed_follows;