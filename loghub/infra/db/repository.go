package db

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/hugovantighem/backend-assignments/loghub/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	connStr string
}

func NewRepository(connStr string) domain.Repository {
	return Repository{
		connStr: connStr,
	}
}

func (x Repository) Save(ctx context.Context, actor string, action string, occuredAt time.Time) error {
	dbpool, err := pgxpool.New(ctx, x.connStr) // TODO centralize conn
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	_, err = dbpool.Exec(ctx, "insert into LOG_ENTRY(actor,action,occuredAt) values ($1,$2,$3)",
		actor, action, occuredAt)
	if err != nil {
		return fmt.Errorf("repository save error: %w", err)
	}
	return nil
}
