package main

import(
	"encoding/json"
	"net/http"

	"github.com/rigofekete/chirpy/internal/auth"
	"github.com/rigofekete/chirpy/internal/database"
)


func (cfg *apiConfig) handlerUsersUpdate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Password string		`json:"password"`
		Email string 		`json:"email"`
	}

	type response struct {
		User
	}
	
	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't find token", err) 
		return
	}

	userID, err := auth.ValidateJWT(token, cfg.jwtSecret) 
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Unauthorized token", err) 
		return
	}
	
	decoder := json.NewDecoder(r.Body)

	params := parameters{}

	err = decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	hash, err := auth.HashPassword(params.Password)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't hash password", err)
		return
	}
	

	userParams := database.UpdateUserByIDParams{
		Email:			params.Email,
		HashedPassword: 	hash,
		ID:			userID,

	}

	userData, err := cfg.db.UpdateUserByID(r.Context(), userParams)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't update user data", err)
		return
	}

	userJSON := response{
		User: User{
			ID:		userData.ID,
			CreatedAt:	userData.CreatedAt,
			UpdatedAt:	userData.UpdatedAt,
			Email:		userData.Email,
		},
	}

	respondWithJSON(w, http.StatusOK, userJSON)
	return
}
