package loglib

import (
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestEmit(t *testing.T) {
	logrus.SetLevel(logrus.DebugLevel)

	err := Emit(LogEntry{Message: "test"})

	assert.NoError(t, err)

}
