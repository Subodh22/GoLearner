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
	APIKey    string    `json:"api_key`
}

type Feed struct {
	ID        uuid.UUID `json:"id"`
	Createdat time.Time `json:"created_at"`
	Updatedat time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	UserID    uuid.UUID `json:"userId`
	Url       string    `json:"url`
}
type Follower struct {
	ID        uuid.UUID `json:"id"`
	Createdat time.Time `json:"created_at"`
	Updatedat time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	UserID    uuid.UUID `json:"userId`
	FeedID    uuid.UUID `json:"feedId`
}

func databaseUserToUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		Createdat: dbUser.Createdat,
		Updatedat: dbUser.Updatedat,
		Name:      dbUser.Name,
		APIKey:    dbUser.ApiKey,
	}

}
func databaseFeedToFeed(dbFeed database.Feed) Feed {
	return Feed{
		ID:        dbFeed.ID,
		Createdat: dbFeed.Createdat,
		Updatedat: dbFeed.Updatedat,
		Name:      dbFeed.Name,
		UserID:    dbFeed.UserID,
		Url:       dbFeed.Url,
	}

}
func databaseFeedToFeeds(dbFeed []database.Feed) []Feed {
	feeds := []Feed{}
	for _, dbFeed := range dbFeed {
		feeds = append(feeds, databaseFeedToFeed(dbFeed))
	}
	return feeds
}
func databaseFeedToUser(dbFeedToUser database.Feedtouser) Follower {
	return Follower{
		ID:        dbFeedToUser.ID,
		Createdat: dbFeedToUser.Createdat,
		Updatedat: dbFeedToUser.Updatedat,
		UserID:    dbFeedToUser.UserID,
		FeedID:    dbFeedToUser.FeedID,
	}

}
