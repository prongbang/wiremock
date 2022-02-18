package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/prongbang/wiremock/pkg/api/home"
	"github.com/prongbang/wiremock/pkg/api/wiremock"
)

type Routers interface {
	Initials(app *fiber.App)
}

type routers struct {
	HomeRoute     home.Router
	WiremockRoute wiremock.Router
}

func (r *routers) Initials(app *fiber.App) {
	r.HomeRoute.Initial(app)
	r.WiremockRoute.Initial(app)
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
