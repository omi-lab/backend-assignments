package broker

import (
	"context"

	"github.com/hugovantighem/backend-assignments/loglib"
)

type LogEmitter struct {
}

func (x LogEmitter) Emit(ctx context.Context, msg string) error {
	entry := loglib.LogEntry{
		Message: msg,
	}
	err := loglib.Emit(entry)
	return err
}
