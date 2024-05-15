package api

import (
	"gom/core/service"
	"gom/core/sys"
	"gom/pod/identity/db"
	"net/http"
)

func User(c service.Context) (account *db.Account, err error) {

	ctx, err := sys.AuthContextFromRequest(c)

	if err != nil {
		return nil, err
	}

	return db.FindAccountByID(ctx.User)
}

func AuthGuard(next service.Handler) service.Handler {
	return func(c service.Context) {

		ctx, err := sys.AuthContextFromRequest(c)

		if err != nil || ctx == nil {
			c.Status(http.StatusUnauthorized)
			return
		}

		next(c)
	}
}
