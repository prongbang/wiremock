package wiremock

import (
	"github.com/gorilla/mux"
	"github.com/prongbang/wiremock/v2/pkg/config"
	"github.com/prongbang/wiremock/v2/pkg/status"
	"os"
)

type Router interface {
	Initial(route *mux.Router)
}

type route struct {
	UseCase UseCase
}

func (r *route) Initial(route *mux.Router) {

	pattern := status.Pattern()

	// Read dir mock
	files, err := os.ReadDir(config.MockPath)
	if err != nil {
		panic(pattern)
	}

	// Read mock directory
	for _, f := range files {
		if f.IsDir() {

			// Get routes from yaml config
			routes := r.UseCase.GetRoutes(f.Name())

			// Register routers
			for rte := range routes.Routers {
				routers := routes.Routers[rte]
				request := routers.Request
				routers.Response.FileName = f.Name()
				handle := NewHandler(r.UseCase, routers)
				route.HandleFunc(request.URL, handle.Handle).Methods(request.Method)
			}
		}
	}
}

func NewRouter(useCase UseCase) Router {
	return &route{
		UseCase: useCase,
	}
}
