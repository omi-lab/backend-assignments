package app

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
)

type UseCase struct {
	emitter LogEmitter
}

func NewUseCase(emitter LogEmitter) *UseCase {
	return &UseCase{emitter: emitter}
}

func (x UseCase) AppendLog(ctx context.Context, actorID string, msg string) error {
	logrus.Debugf("AppendLog: actor=%q,  msg=%q", actorID, msg)

	if len(actorID) == 0 {
		// some validation
		return fmt.Errorf("actorID should not be empty")
	}

	if len(msg) == 0 {
		// some validation
		return fmt.Errorf("message should not be empty")
	}

	err := x.emitter.Emit(ctx, actorID, msg)
	if err != nil {
		return fmt.Errorf("AppendLog emitter.Emit: %w", err)
	}

	return nil
}
