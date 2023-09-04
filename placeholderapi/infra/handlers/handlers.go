package handlers

import (
	"context"
	"net/http"

	"github.com/hugovantighem/backend-assignments/placeholderapi/api"
	"github.com/hugovantighem/backend-assignments/placeholderapi/app"
	"github.com/hugovantighem/backend-assignments/placeholderapi/infra/broker"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type Api struct{}

func (x Api) Ping(ctx echo.Context) error {
	result := api.Pong{Value: "pong"}
	return ctx.JSON(http.StatusOK, result)
}

func (x Api) Accounts(ctx echo.Context, id string) error {
	// parse request
	cmd := &api.AccountCommand{}
	if err := ctx.Bind(cmd); err != nil {
		return err
	}

	// call usecase
	emitter := broker.LogEmitter{} // TODO: use factory as attribute
	uc := app.NewUseCase(emitter)
	err := uc.AppendLog(context.Background(), id, cmd.Msg)

	// response
	if err != nil {
		logrus.Errorf("Accounts error: %v", err)

		errStr := err.Error()
		resp := api.Response{
			Error: &errStr,
		}
		return ctx.JSON(http.StatusOK, resp)
	}

	result := api.Response_Result{}
	result.FromAccountResult(api.AccountResult{Result: "ok"})
	resp := api.Response{
		Result: &result,
	}

	return ctx.JSON(http.StatusOK, resp)
}
