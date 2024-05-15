package api

import (
	"gom/core/service"
)

var Adapters = []string{
	"TWITTER",
	"LINKEDIN",
	"FACEBOOK",
	"GOOGLE",
	"OUTLOOK",
	"YAHOO",
}

func AuthPing(c service.Context) {

}
