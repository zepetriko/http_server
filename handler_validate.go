package main

import (
	"encoding/json"
	"net/http"
)

func handlerValidateChirp(w http.ResponseWriter, r *http.Request) {
	type chirpRequest struct {
		Body string `json:"body"`
	}

	type cleanedResponse struct {
		CleanedBody string `json:"cleaned_body"`
	}

	decoder := json.NewDecoder(r.Body)

	chirp := chirpRequest{}

	err := decoder.Decode(&chirp)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Something went wrong")
		return
	}

	if len(chirp.Body) > 140 {
		respondWithError(w, http.StatusBadRequest, "Chirp is too long")
		return
	}

	cleaned := cleanProfanity(chirp.Body)

	respondWithJSON(w, http.StatusOK, cleanedResponse{
		CleanedBody: cleaned,
	})
}
