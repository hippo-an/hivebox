package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/hippo-an/hivebox/config"
)

func GetVersion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(VersionResponse{Version: config.AppConfig.Application.Version})
}
