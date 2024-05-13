package main

import (
	"edx/core/sys"
	"edx/pod/identity/api"
	"edx/pod/identity/db"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func main() {

	sys.LoadEnv()

	db.Init()

	mux := chi.NewMux()

	mux.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	mux.Post("/", api.Create)
	mux.Post("/auth", api.Auth)
	mux.Post("/auth/ping/{adapter}", api.AuthPing)
	mux.Get("/me", api.AuthGuard(api.Me))
	mux.Put("/me", api.AuthGuard(api.UpdateMe))
	mux.Put("/me/meta", api.AuthGuard(api.UpdateMeta))
	mux.Get("/verify", api.Verify)
	mux.Post("/recover", api.Recover)
	mux.Post("/health", api.Health)

	listen := os.Getenv("LISTEN")

	sys.Logger().Infof("Listening on %s", listen)

	http.ListenAndServe(listen, mux)
}
