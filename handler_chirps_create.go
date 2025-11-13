package main 

import (
	"net/http"
	"encoding/json"
	"strings"
	"time"
	"errors"


	"github.com/google/uuid"
	"github.com/rigofekete/chirpy/internal/database"
	"github.com/rigofekete/chirpy/internal/auth"
)


type Chirp struct {
	ID 		uuid.UUID	`json:"id"`
	CreatedAt 	time.Time	`json:"created_at"`
	UpdatedAt 	time.Time	`json:"updated_at"`
	UserID 		uuid.UUID	`json:"user_id"`
	Body 		string 		`json:"body"`
}

func getCleanedBody(body string, profane_words map[string]struct{}) string {
	wordsList := strings.Split(body, " ")
 
	for i, word := range wordsList {
		wordLow := strings.ToLower(word)
		if _, exits := profane_words[wordLow]; exits {
			wordsList[i] = "****"
		}
	}

	return strings.Join(wordsList, " ")
}

func validateChirp(body string) (string, error) {
	const maxChirpLength = 140
	if len(body) > maxChirpLength {
		return "", errors.New("Chirp is too long") 
	}


	profane_words := map[string]struct{}{
		"kerfuffle": {}, 
		"sharbert": {}, 
		"fornax": {},
	}

	cleaned := getCleanedBody(body, profane_words)
	return cleaned, nil
}

func (cfg *apiConfig) handlerChirpsCreate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body string `json:"body"`
	}

	tokenString, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Error getting token from bearer", err)
		return
	}

	userID, err := auth.ValidateJWT(tokenString, cfg.jwtSecret)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Invalid JWT token", err)
		return
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err = decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	cleaned, err := validateChirp(params.Body)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error(), err)
		return
	}
	
	chirpParams := database.CreateChirpParams{
		Body: cleaned,
		UserID: userID,
	}

	chirp, err := cfg.db.CreateChirp(r.Context(), chirpParams)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error creating chirp", err)
		return 
	}


	chirpResp := Chirp{
		ID: chirp.ID,
		CreatedAt: chirp.CreatedAt,
		UpdatedAt: chirp.UpdatedAt,
		Body:	   chirp.Body,
		UserID:	   chirp.UserID,
	}
	

	respondWithJSON(w, http.StatusCreated, chirpResp)
	return 
}
