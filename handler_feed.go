package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Adosh74/rssagg/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}

	err := decoder.Decode(&params)

	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error parsing JSON: %s", err))
		return
	}
	modifiedUserId := uuid.NullUUID{
		UUID:  user.ID,
		Valid: true,
	}

	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.Url,
		UserID:    modifiedUserId,
	})

	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error creating feed: %v", err))
	}

	responseWithJson(w, 201, databaseFeedToFeed(feed))
}

func (apiCfg *apiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := apiCfg.DB.GetFeeds(r.Context())

	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error getting feeds: %v", err))
	}

	responseWithJson(w, 200, databaseFeedsToFeeds(feeds))
}
