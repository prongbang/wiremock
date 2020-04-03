package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net/http"
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
	├── mock
	│   ├── login
	│   │   └── route.yml
	│   └── user
	│       ├── response
	│       │   └── user.json
	│       └── route.yml
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

			r.HandleFunc(request.URL, func(w http.ResponseWriter, r *http.Request) {
				// Log
				log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)

				// Prepared response
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(response.Status)
				if response.BodyFile != "" {
					bodyFile := fmt.Sprintf("./mock/%s/response/%s", f.Name(), response.BodyFile)
					source, err := ioutil.ReadFile(bodyFile)
					if err != nil {
						_, _ = w.Write([]byte("{}"))
					} else {
						_, _ = w.Write(source)
					}
				} else {
					_, _ = w.Write([]byte(response.Body))
				}

			}).Methods(request.Method)
		}
	}

	started := fmt.Sprintf(`
 -> wiremock server started on %s
`, port)
	fmt.Println(started)

	_ = http.ListenAndServe(port, r)
}
