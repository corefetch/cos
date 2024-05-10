package main

import (
	"cozin/identity/api"
	"cozin/identity/db"
	"cozin/identity/sys"
	"flag"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {

	devmode := flag.Bool("dev", false, "Specify development mode")
	flag.Parse()

	if *devmode {
		godotenv.Load(".env.dev")
	} else {
		godotenv.Load()
	}

	db.Init()

	mux := chi.NewMux()

	mux.Post("/", api.Create)
	mux.Post("/auth", api.Auth)
	mux.Post("/me", api.Me)
	mux.Get("/verify", api.Verify)
	mux.Post("/recover", api.Recover)
	mux.Post("/ping/{adapter}", api.Ping)

	listen := os.Getenv("LISTEN")

	sys.Logger().Infof("Listening on %s", listen)

	http.ListenAndServe(listen, mux)
}
