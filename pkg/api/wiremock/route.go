package wiremock

import (
	"io/ioutil"
	"log"

	"github.com/gorilla/mux"
	"github.com/prongbang/wiremock/pkg/config"
	"github.com/prongbang/wiremock/pkg/status"
	"gopkg.in/yaml.v2"
)

type Route interface {
	Initial(router *mux.Router)
}

type route struct {
	UseCase UseCase
}

func (r *route) Initial(router *mux.Router) {

	pattern := status.Pattern()

	// Read dir mock
	files, err := ioutil.ReadDir(config.MockPath)
	if err != nil {
		panic(pattern)
	}

	// Read mock directory
	for _, f := range files {
		if f.IsDir() {

			// Read yaml config
			source := r.UseCase.ReadSourceRouteYml(f.Name())

			// Unmarshal yaml config
			routes := Routes{}
			err = yaml.Unmarshal(source, &routes)
			if err != nil {
				log.Fatalf("error: %v", err)
			}

			// Register routers
			for route := range routes.Routers {
				routers := routes.Routers[route]
				request := routers.Request
				routers.Response.FileName = f.Name()
				handle := NewHandler(r.UseCase, routers)
				router.HandleFunc(request.URL, handle.Handle).Methods(request.Method)
			}
		}
	}
}

func NewRoute(useCase UseCase) Route {
	return &route{
		UseCase: useCase,
	}
}
