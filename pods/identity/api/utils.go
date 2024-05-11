package api

import (
	"edx/pod/identity/db"
	"edx/pod/identity/sys"
	"net/http"
)

func User(r *http.Request) (account *db.Account, err error) {

	ctx, err := sys.AuthContextFromRequest(r)

	if err != nil {
		return nil, err
	}

	return db.FindAccountByID(ctx.User)
}

func AuthGuard(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		ctx, err := sys.AuthContextFromRequest(r)

		if err != nil || ctx == nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next(w, r)
	}
}
