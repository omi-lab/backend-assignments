package main

import (
	"os"
	"runtime"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/hugovantighem/backend-assignments/loghub/domain"
	"github.com/hugovantighem/backend-assignments/loghub/infra/broker"
	"github.com/hugovantighem/backend-assignments/loghub/infra/db"
	"github.com/sirupsen/logrus"
)

func main() {

	logrus.SetLevel(logrus.DebugLevel)

	brokerUrl := os.Getenv("BROKER_URL")
	brokerChannel := "log_entries"

	dbUrl := os.Getenv("DB_URL")

	consumer, err := broker.NewConsumer(
		brokerUrl,
		brokerChannel,
		func() domain.Repository {
			return db.NewRepository(dbUrl)
		},
	)
	if err != nil {
		logrus.Error(err)
		os.Exit(1)
	}

	err = consumer.Run()
	if err != nil {
		logrus.Error(err)
		os.Exit(1)
	}

	runtime.Goexit() // TODO graceful shutdown
}
