package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

var dbPool *pgxpool.Pool

func InitDB(ctx context.Context, connStr string) (*pgxpool.Pool, error) {
	var err error
	dbPool, err = pgxpool.New(ctx, connStr)
	return dbPool, err
}

func ClosePool() {
	logrus.Debug("closing pool")
	dbPool.Close()
}
