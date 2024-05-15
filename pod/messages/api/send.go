package api

import (
	"bytes"
	"encoding/json"
	"gom/core/sys"
	"net/http"
)

type NatsEmailSend struct {
}

func Send(w http.ResponseWriter, r *http.Request) {

	buf := bytes.NewBufferString("")

	json.NewEncoder(buf).Encode(NatsEmailSend{})

	sys.Nats.Publish("cos.messages.send", buf.Bytes())
}
