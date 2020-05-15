package config

type Config struct {
	Port string
}

const (
	MockResponsePath = "./mock/%s/response/%s"
	MockRouteYmlPath = "./mock/%s/route.yml"
	MockPath         = "./mock"
)
