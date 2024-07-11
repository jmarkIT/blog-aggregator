package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/blog-aggregator/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) createUserHandler(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	new_user_id, err := uuid.NewUUID()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't generate UUID")
		return
	}

	new_user := database.CreateUserParams{
		ID:        new_user_id,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      params.Name,
	}

	new_user_response, err := cfg.DB.CreateUser(r.Context(), new_user)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, new_user_response)
}
