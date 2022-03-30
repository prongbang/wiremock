package wiremock

import (
	"fmt"
	"net/http"

	"github.com/prongbang/wiremock/v2/pkg/core"
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
	// Reading form values
	maxMemory := 32 << 20 // 32Mb
	if err := r.ParseMultipartForm(int64(maxMemory)); err != nil {
		_ = r.ParseForm()
	}

	// Prepared request
	httpHeader := core.BindHeader(h.Routers.Request.Header, r)

	// Prepared response
	if len(h.Routers.Response.Header) == 0 {
		w.Header().Set("Content-Type", "application/json")
	}
	for k, v := range h.Routers.Response.Header {
		w.Header().Set(k, fmt.Sprintf("%v", v))
	}

	// Process cases matching
	if len(h.Routers.Request.Cases) > 0 {

		// Process cases matching
		matching := h.UseCase.CasesMatching(r, h.Routers.Response.FileName, h.Routers.Request.Cases, Parameters{
			ReqHeader: ReqHeader{
				HttpHeader: httpHeader,
				MockHeader: h.Routers.Request.Header,
			},
		})

		// Process response
		if matching.IsMatch {
			response := h.UseCase.GetMockResponse(matching.Case.Response)
			w.WriteHeader(matching.Case.Response.Status)
			_, _ = w.Write(response)
		} else {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write(matching.Result)
		}
		return
	}

	// Prepared request
	body := core.BindBody(h.Routers.Request.Body, r)

	// Process parameter matching
	matching := h.UseCase.ParameterMatching(Parameters{
		ReqHeader: ReqHeader{
			HttpHeader: httpHeader,
			MockHeader: h.Routers.Request.Header,
		},
		ReqBody: ReqBody{
			HttpBody: body,
			MockBody: h.Routers.Request.Body,
		},
	})

	// Prepared response
	if matching.IsMatch {
		response := h.UseCase.GetMockResponse(h.Routers.Response)
		w.WriteHeader(h.Routers.Response.Status)
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
