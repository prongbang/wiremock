package api

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"strings"

	"github.com/prongbang/wiremock/pkg/config"
	"github.com/prongbang/wiremock/pkg/status"
)

type API interface {
	Register(cfg config.Config)
}

type api struct {
	Router Routers
}

func (a *api) Register(cfg config.Config) {
	status.Banner()

	conf := fiber.Config{
		DisableStartupMessage: true,
	}
	app := fiber.New(conf)

	// Middleware
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
		AllowMethods: strings.Join([]string{
			fiber.MethodGet,
			fiber.MethodPost,
			fiber.MethodHead,
			fiber.MethodPut,
			fiber.MethodDelete,
			fiber.MethodPatch,
			fiber.MethodOptions,
			fiber.MethodTrace,
			fiber.MethodConnect,
		}, ","),
	}))

	// Router
	a.Router.Initials(app)

	status.Started(cfg.Port)

	// Listening
	err := app.Listen(fmt.Sprintf(":%s", cfg.Port))
	if err != nil {
		panic(err)
	}
}

// NewAPI provide apis
func NewAPI(router Routers) API {
	return &api{
		Router: router,
	}
}
