package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func queryStatus() {
	response, err := http.Get("https://api.snowbird.com/api/v1/dor/lift-trail-report?alphabetical=true")
	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	responseData, err := io.ReadAll(resoonse.Body)
	if err != nil {
		log.Printf("error parsing response body\n")
		return
	}

	fmt.Println(responseData)
}
