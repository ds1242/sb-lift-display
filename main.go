package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable is not set")
		return
	}

	mux := http.NewServeMux()

	// mux.HandleFunc("GET /api/v1/status", cfg.handlerGetStatus)
}
