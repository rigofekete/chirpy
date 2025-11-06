package main

import (
	"net/http"
	"github.com/google/uuid"
)


func (cfg *apiConfig) handlerChirpsGetAll(w http.ResponseWriter, r *http.Request) {
	dbChirps, err := cfg.db.GetChirps(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't get chirps from DB", err)
		return
	}

	chirpsJSON := []Chirp{}

	for _, dbChirp := range dbChirps {
		chirpsJSON = append(chirpsJSON, Chirp{
			ID: 		dbChirp.ID,
			CreatedAt: 	dbChirp.CreatedAt,
			UpdatedAt: 	dbChirp.UpdatedAt,
			Body:	   	dbChirp.Body,
			UserID:		dbChirp.UserID,	
		})
	}


	respondWithJSON(w, http.StatusOK, chirpsJSON)
}


func (cfg *apiConfig) handlerChirpsGetSingle(w http.ResponseWriter, r *http.Request) {
	chirpID := r.PathValue("chirpID")

	chirpUUID, err := uuid.Parse(chirpID) 
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid chirp ID", err)
		return
	}

	dbChirp, err := cfg.db.GetChirp(r.Context(), chirpUUID) 
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Couldn't get chirp from DB", err)
		return
	}

	chirpJSON := Chirp{
		ID: 		dbChirp.ID,
		CreatedAt: 	dbChirp.CreatedAt,
		UpdatedAt: 	dbChirp.UpdatedAt,
		Body:	   	dbChirp.Body,
		UserID:		dbChirp.UserID,	
	}

	respondWithJSON(w, http.StatusOK, chirpJSON)
}
