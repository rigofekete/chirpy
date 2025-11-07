package main 

import(
	"net/http"
	"encoding/json"

	"github.com/rigofekete/chirpy/internal/auth"
)


func (cfg *apiConfig) handlerUserLogin(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Password string		`json:"password"`
		Email string 		`json:"email"`
	}

	type response struct {
		User
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
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error checking hash password match", err)
		return
	}

	if match != true {
		respondWithError(w, http.StatusUnauthorized, "Incorrect password", nil)
		return
	}


	userResp := response{
		User: User{
			ID: 		user.ID,
			CreatedAt:	user.CreatedAt,	
			UpdatedAt: 	user.UpdatedAt,
			Email:		user.Email,
		},
	}

	respondWithJSON(w, http.StatusOK, userResp)
	return
}
