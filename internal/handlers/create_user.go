package handlers

import (
	"encoding/json"
	"fmt"
	"go_server/internal/database"
	"go_server/internal/repository"
	"net/http"
)

type RequestCreateUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	db := database.DB
	var req RequestCreateUser
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = repository.CreateUser(db, req.Email, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, err = w.Write([]byte(fmt.Sprintf("Account created. Confirmation link sent to %s.", req.Email)))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
