package home

import (
	"github.com/gofiber/fiber/v2"
)

type Router interface {
	Initial(app *fiber.App)
}

type router struct {
	Handle Handler
}

func (r *router) Initial(app *fiber.App) {
	app.Get("/", r.Handle.GetHome)
}

func NewRouter(handle Handler) Router {
	return &router{
		Handle: handle,
	}
}
