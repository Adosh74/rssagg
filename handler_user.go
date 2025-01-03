package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Adosh74/rssagg/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}

	err := decoder.Decode(&params)

	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error parsing JSON: %s", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})

	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error creating user: %s", err))
		return
	}

	responseWithJson(w, 201, databaseUserToUser(user))
}

func handlerGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	responseWithJson(w, 200, databaseUserToUser(user))
}

func (apiCfg *apiConfig) handlerGetPostsForUser(w http.ResponseWriter, r *http.Request, user database.User) {
	posts, err := apiCfg.DB.GetPostsFroUser(r.Context(), database.GetPostsFroUserParams{
		UserID: user.ID,
		Limit:  50,
	})

	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error getting posts: %s", err))
		return
	}

	responseWithJson(w, 200, databasePostsToPosts(posts))
}
