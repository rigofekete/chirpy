package main

import(
	"net/http"
	"fmt"
)


func (cfg *apiConfig) handlerMetrics(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utg-8")
	w.WriteHeader(http.StatusOK)
	s := fmt.Sprintf("Hits: %d", cfg.fileserverHits.Load())
	w.Write([]byte(s))
}


func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.fileserverHits.Add(1)
		next.ServeHTTP(w, r)
	})
}
