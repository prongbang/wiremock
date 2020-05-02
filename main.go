package main

import (
	"flag"

	"github.com/prongbang/wiremock/cmd"
	"github.com/prongbang/wiremock/pkg/config"
)

func main() {
	port := flag.String("port", "8000", "a string port")
	flag.Parse()
	conf := config.Config{
		Port: *port,
	}
	cmd.Run(conf)
}
