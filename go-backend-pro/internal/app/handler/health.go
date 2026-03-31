package handler

import (
	"encoding/json"
	"net/http"

	"github.com/nguyenhuynhnam/go-backend-pro/internal/app/model"
)

func Health(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSON(w, http.StatusMethodNotAllowed, model.APIResponse{
			Status:  "error",
			Message: "method not allowed",
		})
		return
	}

	writeJSON(w, http.StatusOK, model.APIResponse{
		Status:  "ok",
		Message: "service is healthy",
	})
}

func Ping(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSON(w, http.StatusMethodNotAllowed, model.APIResponse{
			Status:  "error",
			Message: "method not allowed",
		})
		return
	}

	writeJSON(w, http.StatusOK, model.APIResponse{
		Status:  "ok",
		Message: "pong",
	})
}

func writeJSON(w http.ResponseWriter, status int, body model.APIResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(body)
}
