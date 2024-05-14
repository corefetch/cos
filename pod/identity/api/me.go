package api

import (
	"cos/core/service"
	"net/http"
	"strings"
)

func Me(c service.Context) {

	account, err := User(c)

	if err != nil {
		c.Status(http.StatusUnauthorized)
		return
	}

	metaFilter := strings.Split(c.Query("meta"), ",")

	if err := c.Write(account.DisplayWithMeta(metaFilter)); err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
}
