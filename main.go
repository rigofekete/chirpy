package main

import (
	"log"
	"net/http"
)


func main() {
	const filepathRoot = "./app"
	const port = "8080"

	mux := http.NewServeMux()

	srv := http.Server{
		Handler: 	mux,
		Addr: 		":" + port,
	}

	mux.Handle("/app/", http.StripPrefix("/app", http.FileServer(http.Dir(filepathRoot))))
	mux.HandleFunc("/healthz", handlerReadiness)

	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
	log.Fatal(srv.ListenAndServe())
}


func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte(http.StatusText(http.StatusOk))
	if err != nil {
		log.Printf("error writing body: %w", err)
	}
}
