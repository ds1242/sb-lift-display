package main

import (
	"sync"
	"time"
)


type Cache struct {
	mu sync.Mutex
	liftStatus *LiftStatus
	fetchedAt time.Time
}

type LiftStatus struct {
	Lift []Lift `json:"lift"`
}

type Lift struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Hours    string `json:"hours"`
	Season   string `json:"season"`
	Status   string `json:"status"`
	Created  int    `json:"created"`
	Updated  int    `json:"updated"`
	Weight   int    `json:"weight"`
}
