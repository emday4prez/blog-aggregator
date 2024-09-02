package main

import (
	"fmt"
	"net/http"
	"strings"
)

func (cfg *apiConfig) handlerGetUserByApiKey(w http.ResponseWriter, r *http.Request) {
    // Extract the Authorization header
    authHeader := r.Header.Get("Authorization")
    
    const prefix = "ApiKey "
    if !strings.HasPrefix(authHeader, prefix) {
        respondWithError(w, http.StatusUnauthorized, "Invalid Authorization Header")
        return
    }

    // Remove the "ApiKey " prefix to get the actual API key
    apiKey := strings.TrimPrefix(authHeader, prefix)

    // Fetch the user by API key
    user, err := cfg.DB.GetUserByApiKey(r.Context(), apiKey)
    if err != nil {
        fmt.Printf("Couldn't find user: %v\n", err.Error())
        respondWithError(w, http.StatusInternalServerError, "Couldn't find user")
        return
    }

    // Respond with user data in JSON format
    respondWithJSON(w, http.StatusOK, databaseUserToUser(user))
}
