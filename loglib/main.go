package loglib

import "fmt"

type LogEntry struct {
	Message string
}

func Emit(entry LogEntry) error {

	return fmt.Errorf("error")
}
