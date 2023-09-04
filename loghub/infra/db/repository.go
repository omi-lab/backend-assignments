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
}

func NewRepository() domain.Repository {
	return Repository{}
}

func (x Repository) Save(ctx context.Context, actor string, action string, occuredAt time.Time) error {
	user := "postgres"
	pw := "changeme"
	host := "localhost:5432"
	db := "postgres"
	dbUrl := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", user, pw, host, db)
	dbpool, err := pgxpool.New(context.Background(), dbUrl) // TODO centralize conn
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	_, err = dbpool.Exec(context.Background(), "insert into LOG_ENTRY(actor,action,occuredAt) values ($1,$2,$3)",
		actor, action, occuredAt)
	if err != nil {
		return fmt.Errorf("Save: %w", err)
	}
	return nil
}
