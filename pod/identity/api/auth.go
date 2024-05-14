package api

import (
	"cos/core/service"
	"cos/core/sys"
	"cos/pod/identity/db"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func Auth(c service.Context) {

	account := &db.Account{}

	if err := c.Read(&account); err != nil {
		c.Error(http.StatusBadRequest, "specify a valid json")
		return
	}

	existent, err := db.FindAccountByLogin(account.Login)

	if err != nil {
		c.Status(http.StatusUnauthorized)
		return
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(existent.Password),
		[]byte(account.Password),
	)

	if err != nil {
		c.Status(http.StatusUnauthorized)
		return
	}

	key, err := sys.CreateAuthKey(sys.AuthContext{
		User:   existent.ID,
		Scope:  sys.ScopeAuth,
		Expire: time.Now().Add(time.Hour),
	})

	if err != nil {
		sys.Logger().Error("Failed to create key for authorization")
		c.Status(http.StatusInternalServerError)
		return
	}

	response := sys.M{
		"access": key,
	}

	if err := c.Write(response); err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
}
