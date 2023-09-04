package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hugovantighem/backend-assignments/loglib"
	"github.com/hugovantighem/backend-assignments/placeholderapi/api"
	"github.com/hugovantighem/backend-assignments/placeholderapi/app"
	"github.com/hugovantighem/backend-assignments/placeholderapi/infra/broker"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type Api struct {
	brokerParams loglib.BrokerParams
}

func NewApi(brokerParams loglib.BrokerParams) (api.ServerInterface, error) {
	if err := brokerParams.Validate(); err != nil {
		return Api{}, fmt.Errorf("wrong broker params: %w", err)
	}
	return Api{
		brokerParams: brokerParams,
	}, nil
}

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
	emitter := broker.NewLogEmitter(x.brokerParams)
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
