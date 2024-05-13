package main

import (
	"edx/core"
	"edx/core/sys"
	"edx/pod/messages/api"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

func main() {

	sys.LoadEnv()

	mux := chi.NewMux()

	mux.Post("/send", api.Send)
	mux.Post("/templates", core.NoOp)
	mux.Put("/templates/{id}", core.NoOp)
	mux.Delete("/template/{id}", core.NoOp)

	listen := os.Getenv("LISTEN")

	sys.Logger().Infof("Listening on %s", listen)

	http.ListenAndServe(listen, mux)
}
