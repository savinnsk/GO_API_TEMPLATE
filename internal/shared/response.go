package shared

import (
	"encoding/json"
	"net/http"

)

func Created(w http.ResponseWriter, body map[string]interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(body)
}