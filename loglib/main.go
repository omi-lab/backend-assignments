package loglib

import (
	"encoding/json"
	"fmt"

	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
)

type LogEntry struct {
	Message string
}

func Emit(entry LogEntry) error {
	logrus.Debugf("Emitting: %s", entry.Message)

	nc, err := nats.Connect("nats://nats:4222") // TODO: use vars / env vars

	if err != nil {
		return fmt.Errorf("connection failed: %w", err)
	}
	defer nc.Close()

	b, err := json.Marshal(entry)

	if err != nil {
		return fmt.Errorf("marshalling entry failed: %w", err)
	}

	err = nc.Publish("foo", b)

	if err != nil {
		return fmt.Errorf("Emit: %w", err)
	}

	return nil
}
