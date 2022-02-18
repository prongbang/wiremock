package wiremock

import (
	"github.com/gofiber/fiber/v2"
	"github.com/prongbang/wiremock/pkg/config"
	"github.com/prongbang/wiremock/pkg/status"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Router interface {
	Initial(app *fiber.App)
}

type route struct {
	UseCase UseCase
}

func (r *route) Initial(app *fiber.App) {

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
				panic(err)
			}

			// Register routers
			for rte := range routes.Routers {
				routers := routes.Routers[rte]
				request := routers.Request
				routers.Response.FileName = f.Name()
				handle := NewHandler(r.UseCase, routers)
				//router.HandleFunc(request.URL, handle.Handle).Methods(request.Method)
				app.Add(request.Method, request.URL, handle.Handle)
			}
		}
	}
}

func NewRouter(useCase UseCase) Router {
	return &route{
		UseCase: useCase,
	}
}
