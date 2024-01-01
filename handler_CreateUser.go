package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"time"

	"github.com/Subodh22/rssFeed/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}

	err := decoder.Decode(&params)

	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error parsing the r.body: %v", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		Createdat: time.Now().UTC(),
		Updatedat: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Couldnt Create user: %v", err))
		return
	}

	responseWithJson(w, 200, databaseUserToUser((user)))
}
func (apiCfg *apiConfig) handleGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	responseWithJson(w, 200, databaseUserToUser((user)))

}
