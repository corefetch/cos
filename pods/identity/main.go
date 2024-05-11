package main

import (
	"edx/pod/identity/api"
	"edx/pod/identity/db"
	"edx/pod/identity/sys"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

func main() {

	sys.LoadEnv()

	db.Init()

	mux := chi.NewMux()

	mux.Post("/", api.Create)
	mux.Post("/auth", api.Auth)
	mux.Get("/me", api.AuthGuard(api.Me))
	mux.Put("/me", api.AuthGuard(api.UpdateMe))
	mux.Put("/me/meta", api.AuthGuard(api.UpdateMeta))
	mux.Get("/verify", api.Verify)
	mux.Post("/recover", api.Recover)
	mux.Post("/ping/{adapter}", api.Ping)

	listen := os.Getenv("LISTEN")

	sys.Logger().Infof("Listening on %s", listen)

	http.ListenAndServe(listen, mux)
}
