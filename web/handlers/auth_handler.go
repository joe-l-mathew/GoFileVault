package handlers

import (
	"encoding/json"
	"net/http"
)

func AuthHandlers(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"name": "joel",
	}

	// Convert data to JSON
	userJSON, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response
	w.Write(userJSON)
}
