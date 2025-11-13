package main 

import(
	"net/http"
	"encoding/json"
	"time"

	"github.com/rigofekete/chirpy/internal/auth"
)


func (cfg *apiConfig) handlerUserLogin(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Password 		string		`json:"password"`
		Email 			string 		`json:"email"`
		ExpiresInSeconds 	int		`json:"expires_in_seconds"`
	}


	type response struct {
		User
		Token		string	`json:"token"`
		RefreshToken	string 	`json:"refresh_token"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return 
	}

	user, err := cfg.db.GetUser(r.Context(), params.Email) 	
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Incorrect email", err) 
		return
	}

	match, err := auth.CheckPasswordHash(params.Password, user.HashedPassword)
	if err != nil || !match {
		respondWithError(w, http.StatusUnauthorized, "Incorrect email or password", err)
		return
	}

	expirationTime := time.Hour 

	if params.ExpiresInSeconds > 0 && params.ExpiresInSeconds < 3600 {
		expirationTime = time.Duration(params.ExpiresInSeconds) * time.Second
	}


	accessToken, err := auth.MakeJWT(user.ID, cfg.jwtSecret, expirationTime) 
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error creating JWT", err)
		return
	}

	userResp := response{
		User: User{
			ID: 		user.ID,
			CreatedAt:	user.CreatedAt,	
			UpdatedAt: 	user.UpdatedAt,
			Email:		user.Email,
		},
		Token:	accessToken,
	}

	respondWithJSON(w, http.StatusOK, userResp)
	return
}
