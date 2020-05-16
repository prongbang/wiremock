package wiremock

type Routes struct {
	Routers map[string]Routers `yaml:"routes"`
}

type Routers struct {
	Request  Request  `yaml:"request"`
	Response Response `yaml:"response"`
}

type Request struct {
	Method string                 `yaml:"method"`
	URL    string                 `yaml:"url"`
	Header map[string]interface{} `yaml:"header"`
	Body   map[string]interface{} `yaml:"body"`
}

type Response struct {
	Status   int    `yaml:"status"`
	Body     string `yaml:"body"`
	BodyFile string `yaml:"body_file"`
	FileName string
}

type Matching struct {
	Result  []byte
	IsMatch bool
}

type Parameters struct {
	HttpReqHeader map[string]interface{}
	MockReqHeader map[string]interface{}
	HttpReqBody   map[string]interface{}
	MockReqBody   map[string]interface{}
}
