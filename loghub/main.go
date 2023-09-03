package main

import (
	"fmt"

	"github.com/hugovantighem/backend-assignments/loglib"
)

func main() {
	entry := loglib.LogEntry{Message: "hello hub"}

	fmt.Println(entry)
}
