package health

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rohanraj7316/hose/configs"
	"github.com/rohanraj7316/logger"
)

func NewRouter(a fiber.Router) {
	// load the configs
	cfg, err := configs.NewServerConfig()
	if err != nil {
		logger.Error(err.Error())
	}

	// initialize your handler
	handler := New(cfg.ProductName, cfg.ModuleName)

	// declare your routes
	a.Get("/", handler.Health)
}
