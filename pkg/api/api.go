package api

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/prongbang/wiremock/v2/pkg/config"
	"github.com/prongbang/wiremock/v2/pkg/status"
	"log"
	"net/http"
)

type API interface {
	Register(cfg config.Config)
}

type api struct {
	Router Routers
}

func (a *api) Register(cfg config.Config) {
	status.Banner()

	r := mux.NewRouter()
	headers := handlers.AllowedHeaders([]string{"*"})
	methods := handlers.AllowedMethods([]string{http.MethodGet, http.MethodPost, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodTrace, http.MethodDelete, http.MethodOptions})
	origins := handlers.AllowedOrigins([]string{"*"})

	// Router
	a.Router.Initials(r)

	status.Started(cfg.Port)

	// Listening
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", cfg.Port), handlers.CORS(headers, methods, origins)(r)))
}

// NewAPI provide apis
func NewAPI(router Routers) API {
	return &api{
		Router: router,
	}
}
