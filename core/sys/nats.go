package sys

import (
	"os"

	"github.com/nats-io/nats.go"
)

var natsConn *nats.Conn

func Events() *nats.Conn {
	return natsConn
}

func NatsLoad() {

	if natsConn != nil {
		return
	}

	natsURL, exists := os.LookupEnv("NATS")

	if !exists {
		Logger().DPanicf("NATS env is required: %s", natsURL)
		return
	}

	Logger().Infof("Connecting to %s", natsURL)

	conn, err := nats.Connect(
		natsURL,
	)

	if err != nil {
		Logger().Errorf("NATS error: %s", err.Error())
		return
	}

	natsConn = conn
}
