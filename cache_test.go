package main

import (
	"testing"
	"time"
)

func TestCreateCache(t *testing.T) {
	cache := &Cache{
		liftStatus: &LiftStatus{
			Lift: []Lift{
				{ID: "1", Name: "Test Lift", Status: "Open"},
			},
		},
		fetchedAt: time.Now(),
	}

	result, _ := cache.GetLifts()

	if result.Lift[0].ID != "1" || result.Lift[0].Name != "Test Lift" {
		t.Errorf("Expected test data, got %+v", result)
	}
}
