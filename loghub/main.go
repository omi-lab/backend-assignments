package main

import (
	"log"

	"github.com/nats-io/nats.go"
)

// func main() {

// 	url := "nats://localhost:4222"
// 	if url == "" {
// 		url = nats.DefaultURL
// 	}

// 	nc, _ := nats.Connect(url)

// 	defer nc.Drain()

// 	nc.Publish("greet.joe", []byte("hello"))

// 	sub, _ := nc.SubscribeSync("greet.*")

// 	msg, _ := sub.NextMsg(10 * time.Millisecond)
// 	fmt.Println("subscribed after a publish...")
// 	fmt.Printf("msg is nil? %v\n", msg == nil)

// 	nc.Publish("greet.joe", []byte("hello"))
// 	nc.Publish("greet.pam", []byte("hello"))

// 	msg, _ = sub.NextMsg(10 * time.Millisecond)
// 	fmt.Printf("msg data: %q on subject %q\n", string(msg.Data), msg.Subject)

// 	msg, _ = sub.NextMsg(10 * time.Millisecond)
// 	fmt.Printf("msg data: %q on subject %q\n", string(msg.Data), msg.Subject)

// 	nc.Publish("greet.bob", []byte("hello"))

// 	msg, _ = sub.NextMsg(10 * time.Millisecond)
// 	fmt.Printf("msg data: %q on subject %q\n", string(msg.Data), msg.Subject)
// }

func main() {

	nc, _ := nats.Connect("nats://localhost:4222")
	log.Println("Connected to " + "nats://localhost:4222")
	_, _ = nc.Subscribe("foo", func(m *nats.Msg) {
		log.Printf("Received a message on %s: %s\n", m.Subject, string(m.Data))
	})

	select {}
}

// func main() {
// 	logrus.SetLevel(logrus.DebugLevel)
// 	nc, err := nats.Connect("nats://localhost:4222")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer nc.Close()

// 	log.Println("Connected to " + "nats://localhost:4222")

// 	_, err = nc.Subscribe("foo", func(msg *nats.Msg) {
// 		logrus.Println("got message")
// 		entry := loglib.LogEntry{}
// 		err := json.Unmarshal(msg.Data, &entry)
// 		if err != nil {
// 			logrus.Errorf("Unmarshal: %v", err)
// 		}
// 		if err == nil {
// 			// Handle the message
// 			logrus.Infof("Received message : %+v\n", entry.Message)
// 		}
// 	})

// 	if err != nil {
// 		logrus.Errorf("Subscribe error: %v", err)
// 		return
// 	}

// 	// Keep the connection alive
// 	runtime.Goexit()

// }
