package main

import (
	"log"
	"os"
	"runtime"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
)

func main() {

	logrus.SetLevel(logrus.DebugLevel)

	brokerUrl := os.Getenv("BROKER_URL")
	brokerChannel := "log_entries"

	nc, _ := nats.Connect(brokerUrl)
	log.Println("Connected to " + brokerUrl)
	_, _ = nc.Subscribe(brokerChannel, func(m *nats.Msg) {
		log.Printf("Received a message on %s: %s\n", m.Subject, string(m.Data))
	})

	runtime.Goexit()
}
