package handlers

import (
	"database/sql"
	"errors"
	"go_server/internal/database"
	"go_server/internal/repository"
	"net/http"
)

func HandleVerifyUser(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")
	if token == "" {
		http.Error(w, "No token provided", http.StatusBadRequest)
	}

	db := database.DB

	err := repository.VerifyUser(db, token)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "No user with provided token found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
