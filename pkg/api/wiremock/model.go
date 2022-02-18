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
	Cases  map[string]Cases       `yaml:"cases"`
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

type CaseMatching struct {
	Result  []byte
	IsMatch bool
	Case    Cases
}

type ReqHeader struct {
	Http map[string]interface{}
	Mock map[string]interface{}
}

type ReqBody struct {
	Http map[string]interface{}
	Mock map[string]interface{}
}

type Parameters struct {
	ReqHeader ReqHeader
	ReqBody   ReqBody
}

type Cases struct {
	Query    map[string]interface{} `yaml:"query"`
	Body     map[string]interface{} `yaml:"body"`
	Response Response               `yaml:"response"`
}
