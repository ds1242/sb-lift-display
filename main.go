package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	/* port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable is not set")
		return
	}

	mux := http.NewServeMux() */

	liftStatus, err := queryLiftStatus()
	if err != nil {
		log.Printf(err.Error())
	}

	fmt.Println(liftStatus)

	fileBytes, err := json.MarshalIndent(liftStatus, "", " ")
	if err != nil {
		log.Printf("unable to marshal data to json file\n")
	}

	err = os.WriteFile("liftStatus.json", fileBytes, 0644)
	if err != nil {
		log.Printf("unable to write to json file\n")
	}

	// mux.HandleFunc("GET /api/v1/status", cfg.handlerGetStatus)
}
