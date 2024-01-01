package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"time"

	"github.com/Subodh22/rssFeed/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handleCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}

	err := decoder.Decode(&params)

	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error parsing the r.body: %v", err))
		return
	}

	feed, err := apiCfg.DB.CreateFeeds(r.Context(), database.CreateFeedsParams{
		ID:        uuid.New(),
		Createdat: time.Now().UTC(),
		Updatedat: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.Url,
		UserID:    user.ID,
	})
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Couldnt Create feed: %v", err))
		return
	}

	responseWithJson(w, 200, databaseFeedToFeed(feed))
}
func (apiCfg *apiConfig) handleGetFeeds(w http.ResponseWriter, r *http.Request) {

	feed, err := apiCfg.DB.GetFeeds(r.Context())
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Couldnt find any feed: %v", err))
		return
	}

	responseWithJson(w, 200, databaseFeedToFeeds(feed))
}
