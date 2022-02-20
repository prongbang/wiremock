package cmd

import (
	"github.com/prongbang/wiremock/v2/pkg/api"
	"github.com/prongbang/wiremock/v2/pkg/api/home"
	"github.com/prongbang/wiremock/v2/pkg/api/wiremock"
	"github.com/prongbang/wiremock/v2/pkg/config"
)

func Run(conf config.Config) {
	homeRoute := home.NewRouter(home.NewHandler(conf))
	wiremockUseCase := wiremock.NewUseCase()
	wiremockRoute := wiremock.NewRouter(wiremockUseCase)
	routers := api.NewRouters(homeRoute, wiremockRoute)
	apis := api.NewAPI(routers)
	apis.Register(conf)
}
