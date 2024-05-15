package api

import (
	"gom/core/service"
	"gom/core/sys"
)

type NatsEmailSend struct {
}

func Send(c service.Context) {
	sys.Logger().Info("sending message")
}
