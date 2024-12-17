package main

import (
	"fmt"
	"net/http"

	"github.com/Adosh74/rssagg/internal/auth"
	"github.com/Adosh74/rssagg/internal/database"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetApiKey(r.Header)

		if err != nil {
			responseWithError(w, 403, fmt.Sprintf("Auth error: %v", err))
			return
		}

		user, err := apiCfg.DB.GetUserByAPIKEY(r.Context(), apiKey)

		if err != nil {
			responseWithError(w, 400, fmt.Sprintf("Con't find user: %v", err))
			return
		}

		handler(w, r, user)
	}
}
