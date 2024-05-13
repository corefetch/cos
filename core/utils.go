package core

import "net/http"

func NoOp(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("noop"))
}
