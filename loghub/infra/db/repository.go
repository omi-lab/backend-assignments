package db

import (
	"context"
	"fmt"
	"time"

	"github.com/hugovantighem/backend-assignments/loghub/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	conn *pgxpool.Conn
}

func NewRepository(conn *pgxpool.Conn) domain.Repository {
	return Repository{
		conn: conn,
	}
}

func (x Repository) Save(ctx context.Context, actor string, action string, occuredAt time.Time) error {
	_, err := x.conn.Exec(ctx, "insert into LOG_ENTRY(actor,action,occuredAt) values ($1,$2,$3)",
		actor, action, occuredAt)
	if err != nil {
		return fmt.Errorf("repository save error: %w", err)
	}
	return nil
}
