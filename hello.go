// Package main provides ...
package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type Handler func(w http.ResponseWriter, r *http.Request) error

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h(w, r); err != nil {
		w.WriteHeader(404)
		w.Write([]byte("bad Request"))
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) error {
	w.Write([]byte("hello world"))
	return nil
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Method("GET", "/", Handler(helloHandler))
	http.ListenAndServe(":3000", r)
}
