package api

import (
	"context"
	"edx/pod/identity/do"
	"edx/pod/identity/sys"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/chi/v5"
	"golang.org/x/oauth2"
)

var Adapters = []string{
	"TWITTER",
	"LINKEDIN",
	"FACEBOOK",
	"GOOGLE",
	"OUTLOOK",
	"YAHOO",
}

func Ping(w http.ResponseWriter, r *http.Request) {

	adapter := strings.ToUpper(chi.URLParam(r, "adapter"))

	index := do.IndexOf[string](Adapters, adapter)

	if index == -1 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("unknown adapter"))
		return
	}

	code := r.URL.Query().Get("code")

	key := os.Getenv(adapter + "_KEY")
	secret := os.Getenv(adapter + "_SECRET")

	conf := &oauth2.Config{
		ClientID:     key,
		ClientSecret: secret,
		RedirectURL:  r.URL.Query().Get("redirect"),
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
			"https://www.googleapis.com/auth/user.birthday.read",
		},
	}

	// Handle the exchange code to initiate a transport.
	tok, err := conf.Exchange(context.Background(), code)
	if err != nil {
		log.Fatal(err)
	}

	client := conf.Client(context.Background(), tok)

	res, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		sys.Logger().Errorf("failed to request user info: %s", err.Error())
		return
	}

	data, err := io.ReadAll(res.Body)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		sys.Logger().Errorf("invalid response for user info: %s", err.Error())
		return
	}

	fmt.Println(string(data))
}
