package main

import (
	"github.com/hugovantighem/backend-assignments/placeholderapi/api"
	"github.com/hugovantighem/backend-assignments/placeholderapi/infra/handlers"
	"github.com/labstack/echo/v4"
)

func main() {

	srv := echo.New()

	loggingApi := handlers.Api{}

	api.RegisterHandlers(srv, loggingApi)

	srv.Logger.Fatal(srv.Start(":8080"))

}
