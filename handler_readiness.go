package main

import (
	"net/http"
	"os"
)

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	appName := os.Getenv("APP_NAME")
	port := os.Getenv("PORT")

	message := appName + " is ready" + " on port " + port
	responseWithJson(w, 200, struct {
		Message string `json:"msg"`
	}{
		Message: message,
	})

}
