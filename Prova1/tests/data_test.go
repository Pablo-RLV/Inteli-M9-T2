package publisher

import (
	publisher "Prova1/publisher"
	"testing"
)

func TestValidationData(t *testing.T) {
	id := "sensor1"
	msg := publisher.MessageSender(id, "geladeira")
	expectedFields := []string{"id", "tipo", "temperatura", "timestamp"}
	for _, field := range expectedFields {
		if _, ok := msg[field]; !ok {
			t.Errorf("field expected: %s", field)
			return
		}
	}
}
