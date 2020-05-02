package cmd

import (
	"github.com/prongbang/wiremock/pkg/api"
	"github.com/prongbang/wiremock/pkg/api/home"
	"github.com/prongbang/wiremock/pkg/api/wiremock"
	"github.com/prongbang/wiremock/pkg/config"
)

func Run(conf config.Config) {
	homeRoute := home.NewRoute(home.NewHandler(conf))
	wiremockRoute := wiremock.NewRoute()
	apis := api.NewAPI(homeRoute, wiremockRoute)
	apis.Register(conf)
}
