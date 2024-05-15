package main

import (
	"fmt"
	"gom/core"
	"gom/core/sys"
	"gom/pod/messages/api"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/nats-io/nats.go"
)

func main() {

	sys.Init()

	mux := chi.NewMux()

	mux.Post("/send", api.Send)
	mux.Post("/templates", core.NoOp)
	mux.Put("/templates/{id}", core.NoOp)
	mux.Delete("/template/{id}", core.NoOp)

	sys.Events().Subscribe("identity.create", func(msg *nats.Msg) {
		fmt.Println(msg.Subject)
		fmt.Println(msg.Data)
	})

	listen := os.Getenv("LISTEN")

	sys.Logger().Infof("Listening on %s", listen)

	http.ListenAndServe(listen, mux)
}
