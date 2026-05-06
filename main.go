package main

import (
	"encoding/json"
	"fmt"
	"io"
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

	http.HandleFunc("GET /api/lifts", handleGetLifts(cache, apiKey))

	http.HandleFunc("GET /api/test-all-open", handleTestAllOpen(cache, apiKey))

	log.Printf("Serving on PORT : %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func handleGetLifts(cache *Cache, apiKey string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
	}
}

func handleTestAllOpen(cache *Cache, apiKey string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-API-Key") != apiKey {
			RespondWithError(w, http.StatusUnauthorized, "unauthorized")
			return
		}

		file, err := os.Open("liftStatus.json")
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, "unable to open lift file")
			return
		}

		fileData, err := io.ReadAll(file)
		if err != nil {
			fmt.Printf("Error reading file body: %s\n", err)
			return
		}

		var practiceStatus LiftStatus
		err = json.Unmarshal(fileData, &practiceStatus)
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, "unable to unmarshall test file")
			return
		}

		defer file.Close()

		RespondWithJSON(w, http.StatusOK, practiceStatus)
	}
}
