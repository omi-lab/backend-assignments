package broker

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hugovantighem/backend-assignments/loghub/app"
	"github.com/hugovantighem/backend-assignments/loghub/domain"
	"github.com/hugovantighem/backend-assignments/loglib"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
)

type Consumer struct {
	brokerUrl   string
	channel     string
	repoFactory func(ctx context.Context) (domain.Repository, *pgxpool.Conn, error)
}

func NewConsumer(
	brokerUrl string,
	channel string,
	repoFactory func(ctx context.Context) (domain.Repository, *pgxpool.Conn, error),
) (Consumer, error) {
	if len(brokerUrl) == 0 {
		return Consumer{}, fmt.Errorf("brokerUrl should not be empty")
	}

	if len(channel) == 0 {
		return Consumer{}, fmt.Errorf("channel should not be empty")
	}

	return Consumer{
		brokerUrl:   brokerUrl,
		channel:     channel,
		repoFactory: repoFactory,
	}, nil
}

// Run subscribe to a channel and delegate incomming messages processing to application layer.
func (x Consumer) Run() error {
	nc, err := nats.Connect(x.brokerUrl)
	if err != nil {
		return fmt.Errorf("canot connect to broker %s: %w", x.brokerUrl, err)
	}

	logrus.Debug("Connected to " + x.brokerUrl)

	_, err = nc.Subscribe(x.channel, func(m *nats.Msg) {
		logrus.Debugf("Received a message on %s: %s\n", m.Subject, string(m.Data))

		entry := loglib.LogEntry{}

		err := json.Unmarshal(m.Data, &entry)
		if err != nil {
			logrus.Error("error while Unmarshal: %w", err)
			return
		}

		ctx := context.Background()

		repo, conn, err := x.repoFactory(ctx)
		if err != nil {
			logrus.Error("error calling repoFactory: %w", err)
			return
		}
		defer conn.Release()

		err = app.HandleEntry(ctx, entry, repo)
		if err != nil {
			logrus.Error("error calling HandleEntry: %w", err)
			return
		}
	})

	if err != nil {
		return fmt.Errorf("canot subscribe to channel %s: %w", x.channel, err)
	}

	return nil
}
