package main 

import(
	"net/http"
	"context"
	"encoding/json"
)


func (cfg *apiConfig) handlerUsers(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Email string `json:"email"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode paramters", err)
		return 
	}

	user, err := cfg.db.CreateUser(context.Background(), params.Email) 	
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create user", err) 
	}

	userResp := users{
		ID: 		user.ID,
		CreatedAt:	user.CreatedAt,	
		UpdatedAt: 	user.UpdatedAt,
		Email:		user.Email,
	}

	respondWithJSON(w, http.StatusOK, userResp)
	return
}
