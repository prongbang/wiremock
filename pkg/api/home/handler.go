package home

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/prongbang/wiremock/pkg/config"
)

type Handler interface {
	GetHome(c *fiber.Ctx) error
}

type handler struct {
	Cfg config.Config
}

func (h *handler) GetHome(c *fiber.Ctx) error {
	return c.SendString(fmt.Sprintf("Wiremock server started on %s", h.Cfg.Port))
}

func NewHandler(cfg config.Config) Handler {
	return &handler{
		Cfg: cfg,
	}
}
