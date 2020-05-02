package wiremock

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

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
