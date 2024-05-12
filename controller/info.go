package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gaspecian/go-to-do-list/config"
)

type Info struct {
	Name        string `json:"name"`
	Version     string `json:"version"`
	Environment string `json:"environment"`
}

func InfoHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	response := Info{
		Name:        config.LoadEnv().Application.Name,
		Version:     config.LoadEnv().Application.Version,
		Environment: config.LoadEnv().Environment,
	}
	json.NewEncoder(w).Encode(response)
}
