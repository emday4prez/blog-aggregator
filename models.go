package main

import (
	"database/sql"
	"time"

	"github.com/emday4prez/blog-aggregator/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"api_key"`
}

func databaseUserToUser(user database.User) User {
	return User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Name:      user.Name,
		ApiKey:    user.ApiKey,
	}
}

type Feed struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      sql.NullString
	Url       sql.NullString
	UserID    uuid.NullUUID
}

func databaseFeedToFeed(feed database.Feed) Feed {
	return Feed{
		ID:        feed.ID,
		CreatedAt: feed.CreatedAt,
		UpdatedAt: feed.UpdatedAt,
		Name:      feed.Name,
		Url:       feed.Url,
		UserID:    feed.UserID,
	}
}

type FeedFollow struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time	
	UserID    uuid.NullUUID
	FeedID    uuid.NullUUID
}
