package api

import (
	"cos/core/service"
	"cos/core/sys"
	"cos/pod/identity/db"
	"cos/pod/identity/do"
	"net/http"
)

func Create(c service.Context) {

	account := &db.Account{}

	if err := c.Read(&account); err != nil {
		c.Error(http.StatusBadRequest, "spcify a valid json")
		return
	}

	sys.Logger().Infof("Creating account for login %s", account.Login)

	if err := do.CreateAccount(account); err != nil {
		sys.Logger().Errorf("error creating account: ", err.Error())
		c.Error(http.StatusBadRequest, err)
		return
	}

	if err := c.Write(account.Display()); err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
}
