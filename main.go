package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/emday4prez/blog-aggregator/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct{
	DB *database.Queries
}

func errorHandler(w http.ResponseWriter, r *http.Request){
	respondWithError(w, 500, "Internal Server Error")
}

func healthzHandler(w http.ResponseWriter, r *http.Request){
	 response := map[string]string{"status": "ok"}
	respondWithJSON(w, 200, response)
}


func main(){
		const filepathRoot = "."
	err := godotenv.Load(".env")	
	if  err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable is not set")
	}

	dbURL := os.Getenv("CONN")
if dbURL == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}
	db, err := sql.Open("postgres", dbURL)
	dbQueries := database.New(db)


apiCfg := apiConfig{
	DB: dbQueries,
}

 mux := http.NewServeMux()
	mux.HandleFunc("GET /v1/err", errorHandler )
	mux.HandleFunc("GET /v1/healthz", healthzHandler )
	mux.HandleFunc("POST /v1/users", apiCfg.handlerUsersCreate)

		srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
	log.Fatal(srv.ListenAndServe())
}