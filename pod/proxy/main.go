package main

import (
	"encoding/json"
	"flag"
	"gom/core/sys"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

type Entry struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Addr string `json:"addr"`
}

var entries = make([]Entry, 0)

type Proxy struct{}

func (p *Proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// create new entry on index posting
	if r.Method == "POST" && r.URL.Path == "/" {

		// create the entry to fill up
		entry := Entry{}

		// decode payload into the entry struct
		if err := json.NewDecoder(r.Body).Decode(&entry); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// verify service name is defined
		if entry.Name == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("please specify name"))
			return
		}

		// verify address name is defined
		if entry.Addr == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("please specify addr"))
			return
		}

		// parse address for validation
		if _, err := url.Parse(entry.Addr); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("please specify a valid service addr"))
			return
		}

		// verifty entry does not already exists
		for _, current := range entries {
			if current.Addr == entry.Addr && current.Name == entry.Name {
				w.WriteHeader(http.StatusConflict)
				return
			}
		}

		// append the new entry
		entries = append(entries, entry)
		w.WriteHeader(http.StatusCreated)

		return
	}

	// try to proxy the request to the service

	var service = strings.Split(r.URL.Path, "/")

	var found *Entry

	// find service entry based on /{service}/endpoint
	for _, current := range entries {
		if len(service) >= 2 && current.Name == service[1] {
			found = &current
			break
		}
	}

	// service not found
	if found == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// get the service addr url
	uri, err := url.Parse(found.Addr)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		sys.Logger().Errorf("invalid addr: %s", err.Error())
		return
	}

	// remove service name from request path
	r.URL.Path = strings.Replace(r.URL.Path, "/"+service[1], "", 1)

	// reverse prxy on the service addr
	proxy := httputil.NewSingleHostReverseProxy(uri)
	proxy.ServeHTTP(w, r)

}

func main() {

	port := flag.String("port", "8800", "Specify port to listen on")
	flag.Parse()

	sys.Logger().Infof("Listen on %s", *port)

	http.ListenAndServe(":"+*port, &Proxy{})
}
