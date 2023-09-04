package loglib

import (
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestEmit(t *testing.T) { // TODO remove or make it IT
	logrus.SetLevel(logrus.DebugLevel)

	err := Emit(LogEntry{Action: "test"}, BrokerParams{})

	assert.NoError(t, err)

}
