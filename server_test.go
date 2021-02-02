package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"testing"
)

func TestListenMux(t *testing.T) {
	config := LoadEnv(".")

	r := mux.NewRouter()
	r.HandleFunc("/{id}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello world")
	})

	ListenMux(config, r)
}
