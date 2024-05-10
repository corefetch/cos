package api

import (
	"corefetch/identity/db"
	"corefetch/identity/do"
	"encoding/json"
	"net/http"
)

func Create(w http.ResponseWriter, r *http.Request) {

	account := &db.Account{}

	if err := json.NewDecoder(r.Body).Decode(account); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("specify a valid json"))
		return
	}

	if err := do.CreateAccount(account); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	if err := json.NewEncoder(w).Encode(account.Display()); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
