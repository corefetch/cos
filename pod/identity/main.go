package main

import (
	"cos/core/service"
	"cos/core/sys"
	"cos/pod/identity/api"
	"cos/pod/identity/db"
	"net/http"
	"os"
)

func main() {

	sys.Init()
	db.Init()

	srv := service.New("identity", "0.0.0")

	srv.Post("/", api.Create)
	srv.Post("/auth", api.Auth)
	srv.Post("/auth/ping/{adapter}", api.AuthPing)
	srv.Get("/me", api.AuthGuard(api.Me))
	srv.Put("/me", api.AuthGuard(api.UpdateMe))
	srv.Put("/me/meta", api.AuthGuard(api.UpdateMeta))
	srv.Get("/verify", api.Verify)
	srv.Post("/recover", api.Recover)
	srv.Post("/health", api.Health)

	listen := os.Getenv("LISTEN")

	sys.Logger().Infof("Listening on %s", listen)

	http.ListenAndServe(listen, srv)
}
