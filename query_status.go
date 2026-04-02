package main

import (
	"encoding/json"
	"io"
	"net/http"
)

func queryLiftStatus() (*LiftStatus, error) {
	response, err := http.Get("https://api.snowbird.com/api/v1/dor/lift-trail-report?alphabetical=true")
	if err != nil {
		return nil, err
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var status LiftStatus

	err = json.Unmarshal(responseData, &status)
	if err != nil {
		return nil, err
	}

	return &status, nil
}
