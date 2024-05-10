package api

import (
	"cozin/identity/db"
	"cozin/identity/sys"
	"encoding/json"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func Auth(w http.ResponseWriter, r *http.Request) {

	account := &db.Account{}

	if err := json.NewDecoder(r.Body).Decode(account); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("specify a valid json"))
		return
	}

	existent, err := db.FindAccountByLogin(account.Login)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(existent.Password),
		[]byte(account.Password),
	)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	key, err := sys.CreateAuthKey(sys.AuthContext{
		User:   existent.ID,
		Scope:  sys.ScopeAuth,
		Expire: time.Now().Add(time.Hour),
	})

	if err != nil {
		sys.Logger().Error("Failed to create key for authorization")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response := sys.M{
		"access": key,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
