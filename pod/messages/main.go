package main

import (
	"fmt"
	"gom/core/service"
	"gom/core/sys"
	"gom/pod/messages/api"
	"net/http"
	"os"

	"github.com/nats-io/nats.go"
)

type In struct{}
type Out struct{}

func main() {

	sys.Init()

	srv := service.New("messages", "0.0.0")

	srv.Post("/send", api.Send)

	sys.Events().Subscribe("identity.create", func(msg *nats.Msg) {
		fmt.Println(msg.Subject)
		fmt.Println(msg.Data)
	})

	listen := os.Getenv("LISTEN")

	sys.Logger().Infof("Listening on %s", listen)

	http.ListenAndServe(listen, srv)
}
