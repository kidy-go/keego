package main

import (
	"fmt"
	"github.com/spf13/cast"
	"net/http"
)

func ListenMux(config *Config, m http.Handler) error {
	listen := cast.ToString(config.Get("server_listen", ":12138"))
	s := &http.Server{Addr: listen, Handler: m}

	fmt.Println("Listen", listen)

	if cast.ToBool(config.Get("server_ssl")) == true &&
		config.Has("server_ssl_cert") &&
		config.Has("server_ssl_key") {

		return s.ListenAndServeTLS(
			cast.ToString(config.Get("server_ssl_cert")),
			cast.ToString(config.Get("server_ssl_key")),
		)
	} else {
		return s.ListenAndServe()
	}
}
