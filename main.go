package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/yaml.v2"
)

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

func main() {
	port := ":8000"

	// http://patorjk.com/software/taag/#p=display&f=Small%20Slant&t=Wiremock
	banner := `
  _      ___                        __  
 | | /| / (_)______ __ _  ___  ____/ /__
 | |/ |/ / / __/ -_)  ' \/ _ \/ __/  '_/
 |__/|__/_/_/  \__/_/_/_/\___/\__/_/\_\
`
	fmt.Println(banner)

	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprint(w, "Wiremock server started on "+port)
	}).Methods("GET")

	pattern := `
	Wiremock require pattern: 
	project
	└── mock
	   ├── login
	   │   └── route.yml
	   └── user
	       ├── response
	       │   └── user.json
	       └── route.yml

	Please back to root project.
`

	// Read dir mock
	files, err := ioutil.ReadDir("./mock")
	if err != nil {
		panic(pattern)
	}

	// Read mock directory
	for _, f := range files {

		// Read yaml config
		filename := fmt.Sprintf("./mock/%s/route.yml", f.Name())
		source, err := ioutil.ReadFile(filename)
		if err != nil {
			panic(pattern)
		}

		// Unmarshal yaml config
		routes := Routes{}
		err = yaml.Unmarshal(source, &routes)
		if err != nil {
			log.Fatalf("error: %v", err)
		}

		// Register routers
		for route := range routes.Routers {
			request := routes.Routers[route].Request
			response := routes.Routers[route].Response
			response.FileName = f.Name()
			handle := NewHandler(response)
			r.HandleFunc(request.URL, handle.Handle).Methods(request.Method)
		}
	}

	started := fmt.Sprintf(`
 -> wiremock server started on %s
`, port)
	fmt.Println(started)

	_ = http.ListenAndServe(port, r)
}

// Handler is a model for handler router
type Handler interface {
	Handle(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	Resp Response
}

func (h *handler) Handle(w http.ResponseWriter, r *http.Request) {
	// Log
	log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)

	// Prepared response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(h.Resp.Status)
	if h.Resp.BodyFile != "" {
		bodyFile := fmt.Sprintf("./mock/%s/response/%s", h.Resp.FileName, h.Resp.BodyFile)
		source, err := ioutil.ReadFile(bodyFile)
		if err != nil {
			log.Printf("%s %s\n", r.RemoteAddr, err)
			_, _ = w.Write([]byte("{}"))
		} else {
			_, _ = w.Write(source)
		}
	} else {
		_, _ = w.Write([]byte(h.Resp.Body))
	}
}

// NewHandler a instance
func NewHandler(resp Response) Handler {
	return &handler{
		Resp: resp,
	}
}
