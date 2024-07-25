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

func usersPostHandler(w http.ResponseWriter, r *http.Request){

}

func main(){
		const filepathRoot = "."
	err := godotenv.Load(".env")	
	if  err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	dbURL := os.Getenv("CONN")

	db, err := sql.Open("postgres", dbURL)
	dbQueries := database.New(db)


 mux := http.NewServeMux()
	mux.HandleFunc("GET /v1/err", errorHandler )
	mux.HandleFunc("GET /v1/healthz", healthzHandler )
	mux.HandleFunc("POST /v1/users", usersPostHandler)

		srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
	log.Fatal(srv.ListenAndServe())
}

