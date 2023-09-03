package app

import (
	"context"
	"fmt"
)

type UseCase struct {
	emitter LogEmitter
}

func NewUseCase(emitter LogEmitter) *UseCase {
	return &UseCase{emitter: emitter}
}

func (x UseCase) AppendLog(ctx context.Context, msg string) error {
	if len(msg) == 0 {
		// some validation
		return fmt.Errorf("message should not be empty")
	}
	return x.emitter.Emit(ctx, msg)
}
