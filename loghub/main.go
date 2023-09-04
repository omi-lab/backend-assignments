package main

import (
	"log"
	"runtime"

	"github.com/nats-io/nats.go"
)

func main() {

	nc, _ := nats.Connect("nats://localhost:4222")
	log.Println("Connected to " + "nats://localhost:4222")
	_, _ = nc.Subscribe("foo", func(m *nats.Msg) {
		log.Printf("Received a message on %s: %s\n", m.Subject, string(m.Data))
	})

	runtime.Goexit()
}
