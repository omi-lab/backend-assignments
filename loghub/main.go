package main

import (
	"log"
	"runtime"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/nats-io/nats.go"
)

// func init() {
// 	m, err := migrate.New("file://Users/hugo/projects/backend-assignments/migrations", "postgres://localhost:5432/database?sslmode=enable")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	if err := m.Up(); err != nil {
// 		log.Fatal(err)
// 	}
// }

func main() {

	nc, _ := nats.Connect("nats://localhost:4222")
	log.Println("Connected to " + "nats://localhost:4222")
	_, _ = nc.Subscribe("foo", func(m *nats.Msg) {
		log.Printf("Received a message on %s: %s\n", m.Subject, string(m.Data))
	})

	runtime.Goexit()
}
