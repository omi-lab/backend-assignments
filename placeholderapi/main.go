package main

import (
	"os"

	"github.com/hugovantighem/backend-assignments/loglib"
	"github.com/hugovantighem/backend-assignments/placeholderapi/api"
	"github.com/hugovantighem/backend-assignments/placeholderapi/infra/handlers"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func main() {

	logrus.SetLevel(logrus.DebugLevel)

	srv := echo.New()

	brokerUrl := os.Getenv("BROKER_URL")
	brokerParams := loglib.BrokerParams{
		Url:     brokerUrl,
		Channel: "log_entries",
	}

	loggingApi, err := handlers.NewApi(brokerParams)
	if err != nil {
		logrus.Errorf("cannot build api: %v", err)
		return
	}

	api.RegisterHandlers(srv, loggingApi)

	srv.Logger.Fatal(srv.Start(":8080")) // TODO grace full shutdown

}
