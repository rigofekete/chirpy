package main

import (
	"net/http"
)


func (cfg *apiConfig) handlerChirpsGet(w http.ResponseWriter, r *http.Request) {
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
	return	
}
