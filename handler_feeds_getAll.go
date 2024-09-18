package main

import (
	"fmt"
	"net/http"
)

func (cfg *apiConfig) handlerGetAllFeeds(w http.ResponseWriter, r *http.Request) {

	feeds, err := cfg.DB.GetAllFeeds(r.Context())

	if err != nil {
		fmt.Printf("Couldn't get feeds: %v\n", err.Error())
		respondWithError(w, http.StatusInternalServerError, "Error Getting All Feeds")
		return
	}

	var feedsSlice []Feed
	for _, v := range feeds {
		feedsSlice = append(feedsSlice, databaseFeedToFeed(v))
	}

	respondWithJSON(w, http.StatusOK, feedsSlice)
}
