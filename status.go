package main

import (
	"encoding/json"
	"net/http"
)

type statusResponse struct {
	Status string `json:"status"`
}

// AliveHandler is the handler for the /status/alive check
func AliveHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	status := statusResponse{
		Status: "Namer service is alive",
	}

	json.NewEncoder(w).Encode(status)
}

// HealthyHandler is the handler for the /status/heathy check
func HealthyHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	status := statusResponse{
		Status: "Namer service is healthy",
	}

	json.NewEncoder(w).Encode(status)
}
