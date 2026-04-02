package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable is not set")
		return
	}

	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatal("missing key value\n")
		return
	}

	cache := &Cache{}
	cache.GetLifts()

	http.HandleFunc("GET /api/lifts", func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-API-Key") != apiKey {
			RespondWithError(w, http.StatusUnauthorized, "unauthorized")
			return
		}
		lifts, err := cache.GetLifts()
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, "failed to fetch lift status")
			return
		}
		RespondWithJSON(w, http.StatusOK, lifts)
	})

	log.Printf("Serving on PORT : %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
