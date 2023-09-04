package loglib

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
)

// LogEntry represent an audit log entry.
type LogEntry struct {
	Actor     string    // the uuid, username, or API token name of the account responsible for the action
	Action    string    // what has been done
	OccuredAt time.Time // at which time the action occured
}

// Emit publishes the entry to a broker.
func Emit(entry LogEntry) error {
	logrus.Debugf("Emitting: %s", entry.Action)

	nc, err := nats.Connect("nats://nats:4222") // TODO: use vars / env vars

	if err != nil {
		return fmt.Errorf("connection failed: %w", err)
	}
	defer nc.Close()

	b, err := json.Marshal(entry)
	if err != nil {
		return fmt.Errorf("marshalling entry failed: %w", err)
	}

	err = nc.Publish("foo", b) // TODO: use queue instead of pubsub
	if err != nil {
		return fmt.Errorf("Emit: %w", err)
	}

	return nil
}
