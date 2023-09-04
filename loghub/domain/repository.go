package domain

import (
	"context"
	"time"
)

//go:generate mockgen -destination=mocks/mock_Repository.go -package=mocks github.com/hugovantighem/backend-assignments/loghub/domain Repository

type Repository interface {
	Save(ctx context.Context, actor string, action string, occuredAt time.Time) error
}
