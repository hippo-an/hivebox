package handlers

import (
	"encoding/json"
	"net/http"
)

const (
	appVersion = "0.0.1"
)

func GetVersion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(VersionResponse{Version: appVersion})
}
