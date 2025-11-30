package handlers

import (
	"encoding/json"
	"go_server/internal/database"
	"go_server/internal/repository"
	"net/http"
	"time"
)

type Time struct {
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
}

func HandleCreateEvent(w http.ResponseWriter, r *http.Request) {
	db := database.DB
	var req Time
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = repository.CreateEvent(db, req.Start, req.End)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
