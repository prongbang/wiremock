package home

import (
	"github.com/gorilla/mux"
)

type Router interface {
	Initial(route *mux.Router)
}

type router struct {
	Handle Handler
}

func (r *router) Initial(route *mux.Router) {
	route.HandleFunc("/", r.Handle.GetHome).Methods("GET")
}

func NewRouter(handle Handler) Router {
	return &router{
		Handle: handle,
	}
}
