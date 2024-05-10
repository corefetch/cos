package api

import (
	"cozin/identity/db"
	"cozin/identity/sys"
	"encoding/json"
	"net/http"
	"strings"
)

func Me(w http.ResponseWriter, r *http.Request) {

	ctx, err := sys.AuthContextFromRequest(r)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	account, err := db.FindAccountByID(ctx.User)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	metaFilter := strings.Split(r.URL.Query().Get("meta"), ",")

	if err := json.NewEncoder(w).Encode(account.DisplayWithMeta(metaFilter)); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
