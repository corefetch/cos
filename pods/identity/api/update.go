package api

import (
	"edx/pod/identity/sys"
	"encoding/json"
	"net/http"
)

func UpdateMeta(w http.ResponseWriter, r *http.Request) {

	meta := make(map[string]string)

	if err := json.NewDecoder(r.Body).Decode(&meta); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	account, err := User(r)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	account.Meta = meta

	if err := account.Save(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		sys.Logger().Error("failed to update meta: ", err)
		return
	}
}
