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
	
	cache := &Cache{}
	cache.GetLifts()
	
	http.HandleFunc("GET /api/lifts", func(w http.ResponseWriter, r *http.Request) {
		lifts, err := cache.GetLifts()
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, "failed to fetch lift status")
			return
		}
		RespondWithJSON(w, http.StatusOK, lifts)
	})

	log.Printf("Serving on PORT : %s\n", port)
	log.Fatal(http.ListenAndServe(":" + port, nil))
}
