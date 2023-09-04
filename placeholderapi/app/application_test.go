package app_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/hugovantighem/backend-assignments/placeholderapi/app"
	"github.com/hugovantighem/backend-assignments/placeholderapi/app/mocks"
	"github.com/stretchr/testify/assert"
)

func TestAppendLog(t *testing.T) {
	t.Run("errors", func(t *testing.T) {
		t.Run("invalid message", func(t *testing.T) {
			// GIVEN an invalid message
			msg := ""
			// WHEN AppendLog
			uc := app.NewUseCase(nil)
			err := uc.AppendLog(context.Background(), uuid.NewString(), msg)
			// THEN an error is returned
			assert.Error(t, err)
		})
		t.Run("emitterError", func(t *testing.T) {
			// setup
			ctrl := gomock.NewController(t)
			emitter := mocks.NewMockLogEmitter(ctrl)
			emitter.EXPECT().Emit(gomock.Any(), gomock.Any(), gomock.Any()).
				Return(fmt.Errorf("error"))
			// GIVEN a valid message
			msg := "foo"
			// WHEN AppendLog
			uc := app.NewUseCase(emitter)
			err := uc.AppendLog(context.Background(), uuid.NewString(), msg)
			// THEN an error is returned
			assert.Error(t, err)
		})
	})
	t.Run("success", func(t *testing.T) {
		// setup
		ctrl := gomock.NewController(t)
		emitter := mocks.NewMockLogEmitter(ctrl)
		emitter.EXPECT().Emit(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(nil)
		// GIVEN a valid message
		msg := "foo"
		// WHEN AppendLog
		uc := app.NewUseCase(emitter)
		err := uc.AppendLog(context.Background(), uuid.NewString(), msg)
		// THEN no error is returned
		assert.NoError(t, err)
	})
}
