package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

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
 mux := http.NewServeMux()
	mux.HandleFunc("GET /v1/err", errorHandler )
	mux.HandleFunc("GET /v1/healthz", healthzHandler )

		srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
	log.Fatal(srv.ListenAndServe())
}

