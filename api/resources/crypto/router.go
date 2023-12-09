package cr

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rohanraj7316/hose/utils/hose"
)

func NewRouter(a fiber.Router) {
	c := hose.New()

	model := NewModel(c)

	handler := NewHandler(model)

	a.Get("/secret-key", handler.GetSecret)
	a.Post("/encrypt", handler.Encrypt)
	a.Post("/decrypt", handler.Decrypt)
}
