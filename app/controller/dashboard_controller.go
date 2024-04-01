package controller

import (
	"encoding/json"
	"net/http"
)

func DahshboardHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "This is protected area!, should have a token to access it.",
	})
}