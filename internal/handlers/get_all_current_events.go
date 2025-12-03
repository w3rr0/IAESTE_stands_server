package handlers

import (
	"encoding/json"
	"go_server/internal/database"
	"go_server/internal/repository"
	"net/http"
)

type RequestGetAllCurrentEvents struct {
	UserId int `json:"user_id"`
}

func HandleGetAllCurrentEvents(w http.ResponseWriter, r *http.Request) {
	db := database.DB
	var req RequestGetAllCurrentEvents
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	eventsList, err := repository.GetAllCurrentEvents(db, req.UserId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(eventsList)
}
