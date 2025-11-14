package main 

import(
	"net/http"
	"time"

	"github.com/rigofekete/chirpy/internal/auth"
)


func (cfg *apiConfig) handlerRefreshToken(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Token string 	`json:"token"`
	}
	
	refreshToken, err := auth.GetBearerToken(r.Header) 

	user, err := cfg.db.GetUserFromRefreshToken(r.Context(), refreshToken)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Error getting user from refresh token", err)
		return
	}

	newJWT, err := auth.MakeJWT(user.ID, cfg.jwtSecret, time.Hour)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error creating JWT", err)
		return
	}

	params := parameters{
		Token: newJWT,
	}

	respondWithJSON(w, http.StatusOK, params)
	return
}



func (cfg *apiConfig) handlerRevokeToken(w http.ResponseWriter, r *http.Request) {
	refreshToken, err := auth.GetBearerToken(r.Header) 

	_, err = cfg.db.RevokeRefreshToken(r.Context(), refreshToken)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Token doesn't exist or is expired/revoked", err) 
		return
	}

	respondWithJSON(w, http.StatusNoContent, nil)
	return 
}
