package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/emday4prez/blog-aggregator/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	feed, err := cfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      sql.NullString{String: params.Name, Valid: true},
		Url:       sql.NullString{String: params.Url, Valid: true},
		UserID:    uuid.NullUUID{UUID: user.ID, Valid: true},
	})
	if err != nil {
		fmt.Printf("Couldn't create feed: %v\n", err.Error())
		respondWithError(w, http.StatusInternalServerError, "Couldn't create feed")
		return
	}
	respondWithJSON(w, http.StatusOK, databaseFeedToFeed(feed))
}
