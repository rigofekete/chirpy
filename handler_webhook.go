package main

import(
	"net/http"
	"encoding/json"

	"github.com/google/uuid"
)


func (cfg *apiConfig) handlerWebhook(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Event 	string	`json:"event"`
		Data	struct {
			UserID	string	`json:"user_id"`
		}	`json:"data"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	if params.Event != "user.upgraded" {
		respondWithJSON(w, http.StatusNoContent, nil)
		return
	}


	userID, err := uuid.Parse(params.Data.UserID)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid user ID", err)
		return
	}

	_, err = cfg.db.UpdateUserChirpyRed(r.Context(), userID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Couldn't update user", err)
		return
	}

	respondWithJSON(w, http.StatusNoContent, nil)
	return
}
