package main

import (
	"fmt"

	"github.com/hugovantighem/backend-assignments/placeholderapi/api"
	"github.com/hugovantighem/backend-assignments/placeholderapi/infra/handlers"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func main() {

	fmt.Println("set log level")
	logrus.SetLevel(logrus.DebugLevel)

	srv := echo.New()

	loggingApi := handlers.Api{}

	api.RegisterHandlers(srv, loggingApi)

	srv.Logger.Fatal(srv.Start(":8080"))

}
