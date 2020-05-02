package wiremock

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/gorilla/mux"
	"github.com/prongbang/wiremock/pkg/status"
	"gopkg.in/yaml.v2"
)

type Route interface {
	Initial(router *mux.Router)
}

type route struct {
}

func (r *route) Initial(router *mux.Router) {

	pattern := status.Pattern()

	// Read dir mock
	files, err := ioutil.ReadDir("./mock")
	if err != nil {
		panic(pattern)
	}

	// Read mock directory
	for _, f := range files {

		// Read yaml config
		filename := fmt.Sprintf("./mock/%s/route.yml", f.Name())
		source, err := ioutil.ReadFile(filename)
		if err != nil {
			panic(pattern)
		}

		// Unmarshal yaml config
		routes := Routes{}
		err = yaml.Unmarshal(source, &routes)
		if err != nil {
			log.Fatalf("error: %v", err)
		}

		// Register routers
		for route := range routes.Routers {
			request := routes.Routers[route].Request
			response := routes.Routers[route].Response
			response.FileName = f.Name()
			handle := NewHandler(response)
			router.HandleFunc(request.URL, handle.Handle).Methods(request.Method)
		}
	}
}

func NewRoute() Route {
	return &route{}
}
