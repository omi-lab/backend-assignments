package app

import "context"

//go:generate mockgen -destination=mocks/mock_LogEmitter.go -package=mocks github.com/hugovantighem/backend-assignments/placeholderapi/app LogEmitter

// LogEmitter defines method to publish message to a broker.
type LogEmitter interface {
	// Emit publishes a message to a brocker.
	Emit(ctx context.Context, actor string, msg string) error
}
