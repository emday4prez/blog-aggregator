package main

import (
	"fmt"
	"net/http"

	"github.com/emday4prez/blog-aggregator/internal/auth"
)

func (cfg *apiConfig) handlerUsersGet(w http.ResponseWriter, r *http.Request) {
	apiKey, err := auth.GetAPIKey(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, fmt.Sprintf("Could not get by user api key: %s", err) )
		return
	}

	user, err := cfg.DB.GetUserByApiKey(r.Context(), apiKey) 
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Couldn't get user")
		return
	}

	respondWithJSON(w, http.StatusOK, databaseUserToUser(user))
}
