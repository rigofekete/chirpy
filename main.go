package main

import (
	"log"
	"net/http"
	"sync/atomic"
	"database/sql"
	"time"
	"os"

	_ "github.com/lib/pq"
	"github.com/joho/godotenv"
	"github.com/rigofekete/chirpy/internal/database"
	"github.com/google/uuid"
)

type apiConfig struct {
	fileserverHits atomic.Int32
	db *database.Queries
}

type users struct {
	ID        uuid.UUID 	`json:"id"`
	CreatedAt time.Time	`json:"created_at"`
	UpdatedAt time.Time	`json:"updated_at"`
	Email 	  string	`json:"email"`
}


func main() {
	const filepathRoot = "./app"
	const port = "8080"

	godotenv.Load()
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL must be set")
	}


	dbConn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Error opening database: %w", err)
	}
	dbQueries := database.New(dbConn)

	apiCfg := apiConfig{
		fileserverHits: atomic.Int32{},
		db:		dbQueries,
	}

	mux := http.NewServeMux()

	srv := http.Server{
		Handler: 	mux,
		Addr: 		":" + port,
	}

	mux.Handle("/app/", apiCfg.middlewareMetricsInc(http.StripPrefix("/app", http.FileServer(http.Dir(filepathRoot)))))
	mux.HandleFunc("GET /api/healthz", handlerReadiness)
	mux.HandleFunc("GET /admin/metrics", apiCfg.handlerMetrics)
	mux.HandleFunc("POST /admin/reset", apiCfg.handlerReset)
	mux.HandleFunc("POST /api/validate_chirp", handlerValidateChirp)
	mux.HandleFunc("POST /api/users", apiCfg.handlerUsers)
	
	
	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
	log.Fatal(srv.ListenAndServe())
}










