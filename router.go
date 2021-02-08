package keego

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"reflect"
)

type Router struct {
	router *mux.Router
	app    *Application
}

type RouteHandler struct {
	App      *Application
	Handler  interface{}
	ValueOf  reflect.Value
	TypeOf   reflect.Type
	Request  *http.Request
	Response http.ResponseWriter
}

func NewRouter(app *Application) Router {
	return Router{
		router: mux.NewRouter(),
		app:    app,
	}
}

func (r Router) Router() *mux.Router {
	return r.router
}

func (r Router) GetApp() *Application {
	return r.app
}

func (r Router) handler(resWrite http.ResponseWriter, request *http.Request, handler interface{}) RouteHandler {
	typ := reflect.TypeOf(handler)
	val := reflect.ValueOf(handler)
	ret := RouteHandler{
		App:      r.app,
		Handler:  handler,
		ValueOf:  val,
		TypeOf:   typ,
		Request:  request,
		Response: resWrite,
	}

	vars := mux.Vars(request)
	fmt.Println(vars)

	switch typ.Kind() {
	case reflect.Func:
		fmt.Println("FUNC:", typ.NumIn(), ":", typ.NumOut())
	case reflect.Ptr:
		fmt.Println("Controller:", typ.NumMethod())
	case reflect.String:
	default:
	}
	return ret
}

func (r Router) controller(route RouteHandler) {
	if reflect.Ptr != route.TypeOf.Kind() {
		return
	}

	typ := route.TypeOf

	for i := 0; i < typ.NumMethod(); i++ {
		m := typ.Method(i)
		fmt.Println(" > ", m.Name, ":", m)
	}
}

func (r Router) dispatch(route RouteHandler) {
	var (
		in  []reflect.Type
		out []reflect.Type
	)

	typ := route.TypeOf
	for i := 0; i < typ.NumIn(); i++ {
		in = append(in, typ.In(i))
	}

	for i := 0; i < typ.NumOut(); i++ {
		out = append(out, typ.Out(i))
	}
}

func (r Router) Handle(path string, handler interface{}) {
	r.Router().HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
		r.handler(w, req, handler)
		fmt.Fprintf(w, "hello kee")
	})
}

func (r Router) Get(path string, handler interface{}) {
	r.Router().HandleFunc(path, func(response http.ResponseWriter, request *http.Request) {
		r.handler(response, request, handler)
		fmt.Fprintf(response, "hello kee")
		//handler(request)
	}).Methods("GET")
}

func (r Router) Group(option map[string]string, handler func(Router)) {
	route := r.Router().NewRoute()
	if prefix, ok := option["prefix"]; ok {
		route.PathPrefix(prefix)
	}

	if host, ok := option["host"]; ok {
		route.Host(host)
	}

	router := Router{
		router: route.Subrouter(),
		app:    r.app,
	}
	handler(router)
}
