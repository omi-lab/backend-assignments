package loglib

import (
	"encoding/json"
	"fmt"

	"github.com/nats-io/nats.go"
)

type LogEntry struct {
	Message string
}

func Emit(entry LogEntry) error {

	nc, err := nats.Connect("nats://nats:4222") // TODO: use vars / env vars

	if err != nil {
		return fmt.Errorf("connection failed: %w", err)
	}

	b, err := json.Marshal(entry)

	if err != nil {
		return fmt.Errorf("marshalling entry failed: %w", err)
	}

	return nc.Publish("foo", b)
}
