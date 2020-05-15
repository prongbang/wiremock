package wiremock

import (
	"log"
	"net/http"

	"github.com/prongbang/wiremock/pkg/api/core"
)

// Handler is a model for handler router
type Handler interface {
	Handle(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	UseCase UseCase
	Routers Routers
}

func (h *handler) Handle(w http.ResponseWriter, r *http.Request) {
	// Log
	log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)

	// Prepared request
	reqt := h.Routers.Request
	data := core.Bind(reqt.Body, r)

	// Process parameter matching
	matching := h.UseCase.ParameterMatching(Parameters{
		HttpReqBody: data,
		MockReqBody: reqt.Body,
	})

	// Prepared response
	w.Header().Set("Content-Type", "application/json")
	if len(reqt.Body) == matching.Count {
		resp := h.Routers.Response
		w.WriteHeader(resp.Status)
		response := h.UseCase.GetMockResponse(resp)
		_, _ = w.Write(response)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write(matching.Result)
	}
}

// NewHandler a instance
func NewHandler(useCase UseCase, routers Routers) Handler {
	return &handler{
		UseCase: useCase,
		Routers: routers,
	}
}
