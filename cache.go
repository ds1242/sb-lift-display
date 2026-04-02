package main

import (
	"time"
)

func (c *Cache) GetLifts() (*LiftStatus, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if time.Since(c.fetchedAt) < time.Hour {
		return c.liftStatus, nil
	}

	liftStatus, err := queryLiftStatus()
	if err != nil {
		return c.liftStatus, nil
	}

	c.liftStatus = liftStatus
	c.fetchedAt = time.Now().UTC()

	return c.liftStatus, nil

}
