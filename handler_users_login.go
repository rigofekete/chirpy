package main 

import(
	"net/http"
	"encoding/json"
	"time"

	"github.com/rigofekete/chirpy/internal/auth"
	"github.com/rigofekete/chirpy/internal/database"
)


func (cfg *apiConfig) handlerUserLogin(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Password 		string		`json:"password"`
		Email 			string 		`json:"email"`
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


	accessToken, err := auth.MakeJWT(user.ID, cfg.jwtSecret, time.Hour) 
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error creating JWT", err)
		return
	}

	refreshToken, _ := auth.MakeRefreshToken()	


	refreshTokenParams := database.CreateRefreshTokenParams{
		Token: 		refreshToken,
		UserID: 	user.ID,
		ExpiresAt: 	time.Now().UTC().Add(time.Hour * 24 * 60),
	}


	_, err = cfg.db.CreateRefreshToken(r.Context(), refreshTokenParams)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error creating refresh token", err)
		return
	}
	
	userResp := response{
		User: User{
			ID: 		user.ID,
			CreatedAt:	user.CreatedAt,	
			UpdatedAt: 	user.UpdatedAt,
			Email:		user.Email,
		},
		Token:		accessToken,
		RefreshToken:	refreshToken,
	}

	respondWithJSON(w, http.StatusOK, userResp)
	return
}
