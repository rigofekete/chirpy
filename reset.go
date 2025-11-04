package main 

import (
	"net/http"
)

func (cfg *apiConfig) handlerReset(w http.ResponseWriter, r *http.Request) {
	if cfg.platform != "dev" {
		// we can ommit this when writing simple non-JSON text for the response, 
		// it will be the default
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")

		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Reset is only allowed in dev environment"))
		return
	}

	cfg.fileserverHits.Store(0)
	// When calling db functions, always pass the http.Request context method
	// since it should be canceled if the client goes away
	err := cfg.db.DeleteUsers(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to reset database: " + err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hits reset to 0 and database reset to initial state"))
	return
}

