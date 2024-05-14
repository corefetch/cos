package main

import (
	"cos/core/sys"
	"encoding/json"
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

func (current Entry) Compare(entry Entry) bool {
	return current.Addr == entry.Addr && current.Name == entry.Name
}

var entries = make([]Entry, 0)

type Proxy struct{}

func (p *Proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" && r.URL.Path == "/" {

		entry := Entry{}

		if err := json.NewDecoder(r.Body).Decode(&entry); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if entry.Name == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("please specify name"))
			return
		}

		if entry.Addr == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("please specify addr"))
			return
		}

		_, err := url.Parse(entry.Addr)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("please specify a valid service addr"))
			return
		}

		for _, current := range entries {
			if current.Compare(entry) {
				w.WriteHeader(http.StatusConflict)
				return
			}
		}

		entries = append(entries, entry)

		w.WriteHeader(http.StatusCreated)

		return
	}

	var service = strings.Split(r.URL.Path, "/")

	var found *Entry

	// find service entry based on /{service}/endpoint
	for _, current := range entries {
		if len(service) >= 2 && current.Name == service[1] {
			found = &current
			break
		}
	}

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
	http.ListenAndServe(":8600", &Proxy{})
}
