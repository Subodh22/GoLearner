package main

import (
	"fmt"
	"net/http"

	"github.com/Subodh22/rssFeed/internal/auth"
	"github.com/Subodh22/rssFeed/internal/database"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *apiConfig) middlewareAuth(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetApiKey(r.Header)
		if err != nil {
			responseWithError(w, 403, fmt.Sprintf("Couldnt find apikey in header: %v", err))
			return
		}
		getUser, err := cfg.DB.GetUserByApiKey(r.Context(), apiKey)

		if err != nil {
			responseWithError(w, 401, fmt.Sprintf("Couldnt find the user: %v", err))
			return
		}
		handler(w, r, getUser)

	}
}
