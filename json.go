package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func responseWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("Responding with 5xx error:", msg)
	}

	type errorResponse struct {
		Error string `json:"error"`
	}

	responseWithJson(w, code, errorResponse{
		Error: msg,
	})
}

func responseWithJson(w http.ResponseWriter, code int, payload interface{}) {
	appName := os.Getenv("APP_NAME")

	dat, err := json.Marshal(struct {
		App  string      `json:"app"`
		Data interface{} `json:"data"`
	}{
		App:  appName,
		Data: payload,
	})

	if err != nil {
		log.Printf("Fail to marshal json response %v", payload)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(code)
	w.Write(dat)
}
