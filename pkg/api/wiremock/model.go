package wiremock

type Routes struct {
	Routers map[string]Routers `yaml:"routes"`
}

type Routers struct {
	Request  Request  `yaml:"request"`
	Response Response `yaml:"response"`
}

type Request struct {
	Method string `yaml:"method"`
	URL    string `yaml:"url"`
}

type Response struct {
	Status   int    `yaml:"status"`
	Body     string `yaml:"body"`
	BodyFile string `yaml:"body_file"`
	FileName string
}
