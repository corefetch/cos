package main

import (
	"edx/core"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Router() chi.Router {

	mux := chi.NewMux()

	mux.Get("/", core.NoOp)

	return mux
}

func main() {
	http.ListenAndServe(":8087", Router())
}
