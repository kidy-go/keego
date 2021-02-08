package mux

import (
	"fmt"
	"net/http"
)

type ServerMux struct {
	Host   string
	Path   string
	routes Maps
}

func NewServerMux() *ServerMux {
	return &ServerMux{
		routes: Maps{},
	}
}

func (s *ServerMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//path = r.URL.EscapedPath()
	//fmt.Println(r.URL.EscapedPath())
	ctx := &MuxContext{
		params: Maps{
			"path": r.URL.EscapedPath(),
		},
		request:  r,
		response: w,
	}

	path := r.URL.Path
	if f, ok := s.routes[path]; ok {
		f.(func(Context))(ctx)
		if ret, ok := ctx.Params()["return"]; ok {
			fmt.Fprintf(w, ret.(string))
			return
		}
	} else {
		fmt.Fprintf(w, "Not Fund Page (404).")
	}
}

func (s *ServerMux) HandleFunc(path string, f func(Context)) {
	s.routes[path] = f
}
