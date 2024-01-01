package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"time"

	"github.com/Subodh22/rssFeed/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handleFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feedId"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}

	err := decoder.Decode(&params)

	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error parsing the r.body: %v", err))
		return
	}

	feedToFollow, err := apiCfg.DB.CreateFeedToUser(r.Context(), database.CreateFeedToUserParams{
		ID:        uuid.New(),
		Createdat: time.Now().UTC(),
		Updatedat: time.Now().UTC(),
		FeedID:    params.FeedID,
		UserID:    user.ID,
	})
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Couldnt Create feed Follow: %v", err))
		return
	}

	responseWithJson(w, 200, databaseFeedToUser(feedToFollow))
}
