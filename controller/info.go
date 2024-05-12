package controller

import (
	"encoding/json"
	"net/http"
)

type Info struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

func InfoHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	response := Info{
		Name:    "to-do-api",
		Version: "1.0.0",
	}
	json.NewEncoder(w).Encode(response)
}
