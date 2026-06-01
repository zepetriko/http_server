package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Email     string    `json:"email"`
}

func (cfg *apiConfig) handlerUsersCreate(w http.ResponseWriter, r *http.Request) {
	type requests struct {
		Email string `json:"email"`
	}

	decoder := json.NewDecoder(r.Body)
	request := requests{}
	err := decoder.Decode(&request)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Something went wrong decoding JSON")
		return
	}

	user, err := cfg.db.CreateUser(r.Context(), request.Email)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "User could not be created")
		return
	}

	respondWithJSON(w, http.StatusCreated, User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Email:     user.Email,
	})

}
