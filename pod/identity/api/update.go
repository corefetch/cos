package api

import (
	"gom/core/service"
	"gom/core/sys"
	"gom/pod/identity/db"
	"gom/pod/identity/do"
	"net/http"
)

func UpdateMe(c service.Context) {

	update := &db.Account{}

	if err := c.Read(update); err != nil {
		c.Error(http.StatusBadRequest, "specify a valid json")
		return
	}

	account, err := User(c)

	if err != nil {
		c.Status(http.StatusUnauthorized)
		return
	}

	account.Names = update.Names

	if update.Password != "" {

		// validate password and encrypt
		password, err := do.CreateSecurePassword(update.Password)

		if err != nil {
			c.Status(http.StatusInternalServerError)
			sys.Logger().Errorf("failed to update password: %s", err.Error())
			return
		}

		// replace password in account
		account.Password = password

	}

	if err := account.Save(); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
}

func UpdateMeta(c service.Context) {

	meta := make(map[string]string)

	if err := c.Read(&meta); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	account, err := User(c)

	if err != nil {
		c.Status(http.StatusUnauthorized)
		return
	}

	account.Meta = meta

	if err := account.Save(); err != nil {
		c.Status(http.StatusInternalServerError)
		sys.Logger().Error("failed to update meta: ", err)
		return
	}
}
