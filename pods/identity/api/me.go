package api

import (
	"encoding/json"
	"net/http"
	"strings"
)

func Me(w http.ResponseWriter, r *http.Request) {

	account, err := User(r)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	metaFilter := strings.Split(r.URL.Query().Get("meta"), ",")

	if err := json.NewEncoder(w).Encode(account.DisplayWithMeta(metaFilter)); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
