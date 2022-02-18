package wiremock

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"

	"github.com/prongbang/wiremock/pkg/api/core"
)

// Handler is a model for handler router
type Handler interface {
	Handle(c *fiber.Ctx) error
}

type handler struct {
	UseCase UseCase
	Routers Routers
}

func (h *handler) Handle(c *fiber.Ctx) error {

	// Prepared request
	httpHeader := core.BindHeader(h.Routers.Request.Header, c)

	// Prepared response
	if len(h.Routers.Response.Header) == 0 {
		c.Response().Header.Set("Content-Type", "application/json")
	}
	for k, v := range h.Routers.Response.Header {
		c.Set(k, fmt.Sprintf("%v", v))
	}

	// Process cases matching
	if len(h.Routers.Request.Cases) > 0 {

		// Process cases matching
		matching := h.UseCase.CasesMatching(c, h.Routers.Response.FileName, h.Routers.Request.Cases, Parameters{
			ReqHeader: ReqHeader{
				HttpHeader: httpHeader,
				MockHeader: h.Routers.Request.Header,
			},
		})

		// Process response
		if matching.IsMatch {
			response := h.UseCase.GetMockResponse(matching.Case.Response)
			return c.Status(matching.Case.Response.Status).SendString(string(response))
		}

		return c.Status(http.StatusBadRequest).SendString(string(matching.Result))
	}

	// Prepared request
	body := core.BindBody(h.Routers.Request.Body, c)

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
		return c.Status(h.Routers.Response.Status).SendString(string(response))
	}

	return c.Status(http.StatusBadRequest).SendString(string(matching.Result))
}

// NewHandler a instance
func NewHandler(useCase UseCase, routers Routers) Handler {
	return &handler{
		UseCase: useCase,
		Routers: routers,
	}
}
