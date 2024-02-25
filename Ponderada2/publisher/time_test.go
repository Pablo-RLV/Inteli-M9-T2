package main

import (
	"testing"
	"time"
)

func TestTimeToPublish(t *testing.T) {
	start := time.Now()
	TimeToPublish()
	elapsed := time.Since(start)
	if elapsed < 2*time.Second {
		t.Errorf("TimeToPublish() should sleep for 2 seconds, but slept for %s", elapsed)
	}
}