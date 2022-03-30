package api

import (
	"github.com/gorilla/mux"
	"github.com/prongbang/wiremock/v2/pkg/api/home"
	"github.com/prongbang/wiremock/v2/pkg/api/wiremock"
)

type Routers interface {
	Initials(route *mux.Router)
}

type routers struct {
	HomeRoute     home.Router
	WiremockRoute wiremock.Router
}

func (r *routers) Initials(route *mux.Router) {
	r.HomeRoute.Initial(route)
	r.WiremockRoute.Initial(route)
}

func NewRouters(
	homeRoute home.Router,
	wiremockRoute wiremock.Router,
) Routers {
	return &routers{
		HomeRoute:     homeRoute,
		WiremockRoute: wiremockRoute,
	}
}
