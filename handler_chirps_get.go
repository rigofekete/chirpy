package main

import (
	"net/http"
	"sort"

	"github.com/google/uuid"
)


func (cfg *apiConfig) handlerChirpsGetAll(w http.ResponseWriter, r *http.Request) {
	dbChirps, err := cfg.db.GetChirps(r.Context()) 
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't get Author ID chirps from DB", err)
		return
	}

	query := "asc"
	queryParam := r.URL.Query().Get("sort")
	if queryParam == "desc" {
		query = "desc"
	}



	authorID := uuid.Nil
	authorIDString := r.URL.Query().Get("author_id")
	if authorIDString != "" {
		authorID, err = uuid.Parse(authorIDString)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid author ID", err)
			return
		}
	} 


	chirps := []Chirp{}

	for _, dbChirp := range dbChirps {
		if authorID != uuid.Nil && authorID != dbChirp.UserID {
			continue
		}

		chirps = append(chirps, Chirp{
			ID: 		dbChirp.ID,
			CreatedAt: 	dbChirp.CreatedAt,
			UpdatedAt: 	dbChirp.UpdatedAt,
			Body:	   	dbChirp.Body,
			UserID:		dbChirp.UserID,	
		})
	}

	sort.Slice(chirps, func(i, j int) bool {
		if query == "desc" {
			return chirps[i].CreatedAt.After(chirps[j].CreatedAt)
		} else {
			return chirps[i].CreatedAt.Before(chirps[j].CreatedAt)
		}
	})


	respondWithJSON(w, http.StatusOK, chirps)
}


func (cfg *apiConfig) handlerChirpsGetSingle(w http.ResponseWriter, r *http.Request) {
	chirpIDString := r.PathValue("chirpID")

	chirpID, err := uuid.Parse(chirpIDString) 
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid chirp ID", err)
		return
	}

	dbChirp, err := cfg.db.GetChirp(r.Context(), chirpID) 
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
