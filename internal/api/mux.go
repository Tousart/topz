package api

import (
	"net/http"
)

type Mux struct {
	Handler map[string]http.HandlerFunc
}

func NewMux() *Mux {
	handler := make(map[string]http.HandlerFunc)
	return &Mux{
		Handler: handler,
	}
}

func (mx *Mux) HandleFunc(method, pattern string, handler http.HandlerFunc) {
	mx.Handler[method+" "+pattern] = handler
}

func (mx *Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if handler, ok := mx.Handler[r.Method+" "+r.URL.Path]; ok {
		handler(w, r)
		return
	}
	w.Write([]byte("method not found"))
}
