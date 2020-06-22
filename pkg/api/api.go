package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/prongbang/wiremock/pkg/api/home"
	"github.com/prongbang/wiremock/pkg/api/wiremock"
	"github.com/prongbang/wiremock/pkg/config"
	"github.com/prongbang/wiremock/pkg/status"
)

type API interface {
	Register(cfg config.Config)
}

type api struct {
	HomeRoute     home.Route
	WiremockRoute wiremock.Route
}

func (a *api) Register(cfg config.Config) {
	status.Banner()

	r := mux.NewRouter()
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization", "X-Platforms"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PATCH", "PUT", "HEAD", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})

	a.HomeRoute.Initial(r)
	a.WiremockRoute.Initial(r)

	status.Started(cfg.Port)

	_ = http.ListenAndServe(fmt.Sprintf(":%s", cfg.Port), handlers.CORS(headers, methods, origins)(r))
}

// NewAPI provide apis
func NewAPI(
	homeRoute home.Route,
	wiremockRoute wiremock.Route,
) API {
	return &api{
		HomeRoute:     homeRoute,
		WiremockRoute: wiremockRoute,
	}
}
