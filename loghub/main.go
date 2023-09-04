package main

import (
	"context"
	"os"
	"runtime"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/hugovantighem/backend-assignments/loghub/domain"
	"github.com/hugovantighem/backend-assignments/loghub/infra/broker"
	"github.com/hugovantighem/backend-assignments/loghub/infra/db"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

func main() {

	ctx := context.Background()

	// config
	logrus.SetLevel(logrus.DebugLevel)

	brokerUrl := os.Getenv("BROKER_URL")
	brokerChannel := "log_entries"

	dbUrl := os.Getenv("DB_URL")

	pool, err := db.InitDB(ctx, dbUrl)
	if err != nil {
		logrus.Error(err)
		os.Exit(1)
	}

	// setup
	consumer, err := broker.NewConsumer(
		brokerUrl,
		brokerChannel,
		func(ctx context.Context) (domain.Repository, *pgxpool.Conn, error) {
			conn, err := pool.Acquire(ctx)
			if err != nil {
				return nil, nil, err
			}

			return db.NewRepository(conn), conn, nil
		},
	)
	if err != nil {
		logrus.Error(err)
		os.Exit(1)
	}

	// consume messages
	err = consumer.Run()
	if err != nil {
		logrus.Error(err)
		os.Exit(1)
	}

	// defer db.ClosePool() // TODO graceful shutdown
	runtime.Goexit() // TODO graceful shutdown
}
