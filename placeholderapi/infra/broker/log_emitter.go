package broker

import (
	"context"
	"fmt"
	"time"

	"github.com/hugovantighem/backend-assignments/loglib"
	"github.com/hugovantighem/backend-assignments/placeholderapi/app"
)

type LogEmitter struct {
	params loglib.BrokerParams
}

func NewLogEmitter(params loglib.BrokerParams) app.LogEmitter {
	return LogEmitter{
		params: params,
	}
}

func (x LogEmitter) Emit(ctx context.Context, actor string, msg string) error {
	entry := loglib.LogEntry{
		Actor:     actor,
		Action:    msg,
		OccuredAt: time.Now(),
	}
	err := loglib.Emit(entry, x.params)
	if err != nil {
		return fmt.Errorf("Emit: %w", err)
	}
	return nil
}
