package home

import "github.com/gorilla/mux"

type Route interface {
	Initial(route *mux.Router)
}

type route struct {
	Handle Handler
}

func (r *route) Initial(route *mux.Router) {
	route.HandleFunc("/", r.Handle.GetHome).Methods("GET")
}

func NewRoute(handle Handler) Route {
	return &route{
		Handle: handle,
	}
}
