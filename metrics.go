package main

import(
	"net/http"
	"fmt"
)


func (cfg *apiConfig) handlerMetrics(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utg-8")
	w.WriteHeader(http.StatusOK)

	s := fmt.Sprintf(`
	<html>
	  <body>
	    <h1>Welcome, Chirpy Admin</h1>
	    <p>Chirpy has been visited %d times!</p>
	  </body>
	</html>`, cfg.fileserverHits.Load())

	w.Write([]byte(s))
}


func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.fileserverHits.Add(1)
		next.ServeHTTP(w, r)
	})
}
