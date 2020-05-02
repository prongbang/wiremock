package home

import (
	"fmt"
	"net/http"

	"github.com/prongbang/wiremock/pkg/config"
)

type Handler interface {
	GetHome(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	Cfg config.Config
}

func (h *handler) GetHome(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprint(w, "Wiremock server started on "+h.Cfg.Port)
}

func NewHandler(cfg config.Config) Handler {
	return &handler{
		Cfg: cfg,
	}
}
