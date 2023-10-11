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
	Query  map[string]interface{} `yaml:"query"`
	Cases  map[string]Cases       `yaml:"cases"`
}

type Response struct {
	Status   int                    `yaml:"status"`
	Header   map[string]interface{} `yaml:"header"`
	Body     string                 `yaml:"body"`
	BodyFile string                 `yaml:"body_file"`
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
	HttpHeader map[string]interface{}
	MockHeader map[string]interface{}
}

type ResHeader struct {
	MockHeader map[string]interface{}
}

type ReqBody struct {
	HttpBody map[string]interface{}
	MockBody map[string]interface{}
}

type ReqQuery struct {
	HttpQuery map[string]interface{}
	MockQuery map[string]interface{}
}

type Parameters struct {
	ReqHeader ReqHeader
	ResHeader ResHeader
	ReqQuery  ReqQuery
	ReqBody   ReqBody
}

type Cases struct {
	Query    map[string]interface{} `yaml:"query"`
	Body     map[string]interface{} `yaml:"body"`
	Response Response               `yaml:"response"`
}
