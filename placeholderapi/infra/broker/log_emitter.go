package broker

import (
	"context"
	"fmt"
	"time"

	"github.com/hugovantighem/backend-assignments/loglib"
)

type LogEmitter struct {
}

func (x LogEmitter) Emit(ctx context.Context, actor string, msg string) error {
	entry := loglib.LogEntry{
		Actor:     actor,
		Action:    msg,
		OccuredAt: time.Now(),
	}
	err := loglib.Emit(entry)
	if err != nil {
		return fmt.Errorf("Emit: %w", err)
	}
	return nil
}
