package app_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/hugovantighem/backend-assignments/loghub/app"
	"github.com/hugovantighem/backend-assignments/loghub/domain/mocks"
	"github.com/hugovantighem/backend-assignments/loglib"
	"github.com/stretchr/testify/assert"
)

func TestHandleEntry(t *testing.T) {
	t.Run("repositoryError", func(t *testing.T) {
		// setup
		ctrl := gomock.NewController(t)
		repo := mocks.NewMockRepository(ctrl)
		repo.EXPECT().Save(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(fmt.Errorf("error"))

		// GIVEN a log entry
		entry := loglib.LogEntry{
			Actor:     "34H4I234U3H424U2U4J",
			Action:    "Create",
			OccuredAt: time.Now(),
		}

		// WHEN handling the enty
		err := app.HandleEntry(context.Background(), entry, repo)

		// THEN an error is raised
		assert.Error(t, err)
	})

	t.Run("success", func(t *testing.T) {
		// setup
		ctrl := gomock.NewController(t)
		repo := mocks.NewMockRepository(ctrl)
		repo.EXPECT().Save(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)

		// GIVEN a log entry
		entry := loglib.LogEntry{
			Actor:     "34H4I234U3H424U2U4J",
			Action:    "Create",
			OccuredAt: time.Now(),
		}

		// WHEN handling the enty
		err := app.HandleEntry(context.Background(), entry, repo)

		// THEN no error is raised
		assert.NoError(t, err)
	})
}
