package main

import (
	"flag"
	"os"

	"github.com/prongbang/wiremock/v2/cmd"
	"github.com/prongbang/wiremock/v2/pkg/config"
)

func main() {
	port := flag.String("port", "8000", "a string port")
	flag.Parse()
	p := *port
	if os.Getenv("PORT") != "" {
		p = os.Getenv("PORT")
	}
	conf := config.Config{
		Port: p,
	}
	cmd.Run(conf)
}
