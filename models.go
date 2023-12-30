package main

import (
	"time"

	"github.com/Subodh22/rssFeed/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Createdat time.Time `json:"created_at"`
	Updatedat time.Time `json:"updated_at"`
	Name      string    `json:"name"`
}

func databaseUserToUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		Createdat: dbUser.Createdat,
		Updatedat: dbUser.Updatedat,
		Name:      dbUser.Name,
	}

}
