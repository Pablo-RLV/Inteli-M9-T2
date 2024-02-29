package publisher

import (
	"testing"
	publisher "Ponderada2/publisher"
)

func TestValidationData(t *testing.T) {
	id := "sensor1"
	msg := publisher.MessageSender(id)
	expectedFields := []string{"CO", "ID", "NH3", "NO2"}
	for _, field := range expectedFields {
		if _, ok := msg[field]; !ok {
			t.Errorf("field expected: %s", field)
			return
		}
	}
}