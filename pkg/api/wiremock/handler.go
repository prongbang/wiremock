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
	header := core.BindHeader(h.Routers.Request.Header, r)
	body := core.BindBody(h.Routers.Request.Body, r)

	// Process parameter matching
	matching := h.UseCase.ParameterMatching(Parameters{
		HttpReqHeader: header,
		MockReqHeader: h.Routers.Request.Header,
		HttpReqBody:   body,
		MockReqBody:   h.Routers.Request.Body,
	})

	// Prepared response
	w.Header().Set("Content-Type", "application/json")
	if matching.IsMatch {
		w.WriteHeader(h.Routers.Response.Status)
		response := h.UseCase.GetMockResponse(h.Routers.Response)
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
