package app

import (
	"context"
	"fmt"

	"github.com/hugovantighem/backend-assignments/loghub/domain"
	"github.com/hugovantighem/backend-assignments/loglib"
)

func HandleEntry(ctx context.Context, entry loglib.LogEntry, repo domain.Repository) error {
	err := repo.Save(context.Background(), entry.Actor, entry.Action, entry.OccuredAt)
	if err != nil {
		return fmt.Errorf("cannot save log entry: %w", err)
	}

	return nil
}
