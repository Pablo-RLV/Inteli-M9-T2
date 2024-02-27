package publisher

import (
	"testing"
	"time"
	publisher "Ponderada2/publisher"
)

func TestTimeToPublish(t *testing.T) {
	start := time.Now()
	publisher.TimeToPublish()
	elapsed := time.Since(start)
	if elapsed < 2*time.Second {
		t.Errorf("TimeToPublish() should sleep for 2 seconds, but slept for %s", elapsed)
	}
}