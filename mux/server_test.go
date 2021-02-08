package mux

import (
	"fmt"
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	s := &http.Server{Addr: ":89"}

	mux := NewServerMux()

	mux.HandleFunc("/test", func(ctx Context) {
		ctx.SetParam("return", "hello world")
		fmt.Println(ctx.Params())
	})

	s.Handler = mux

	s.ListenAndServeTLS("server_cert.crt", "private.key")
}
