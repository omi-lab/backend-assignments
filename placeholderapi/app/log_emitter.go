package app

import "context"

//go:generate mockgen -destination=mocks/mock_LogEmitter.go -package=mocks github.com/hugovantighem/backend-assignments/placeholderapi/app LogEmitter

type LogEmitter interface {
	Emit(ctx context.Context, actor string, msg string) error
}
