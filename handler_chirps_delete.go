package main

import(
	"net/http"

	"github.com/rigofekete/chirpy/internal/auth"
	"github.com/google/uuid"
)


func (cfg *apiConfig) handlerChirpsDelete(w http.ResponseWriter, r *http.Request) {
	chirpIDString := r.PathValue("chirpID")

	chirpID, err := uuid.Parse(chirpIDString)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid chirp ID", err)
		return
	}

	token, err := auth.GetBearerToken(r.Header) 
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Error getting token from bearer", err)
		return
	}

	userID, err := auth.ValidateJWT(token, cfg.jwtSecret) 
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Invalid JWT token", err)
		return
	}

	dbChirp, err := cfg.db.GetChirp(r.Context(), chirpID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Couldn't get chirp from DB", err)
		return
	}

	if userID != dbChirp.UserID {
		respondWithError(w, http.StatusForbidden, "User is not the author of the chirp", err)
		return
	}


	err = cfg.db.DeleteChirp(r.Context(), chirpID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Couldn't get chirp from DB", err)
		return
	}
	
	respondWithJSON(w, http.StatusNoContent, nil)
	return
}
