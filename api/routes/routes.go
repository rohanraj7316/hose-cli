package routes

import (
	"fmt"
	"net/http"

	cr "github.com/rohanraj7316/hose/api/resources/crypto"
	"github.com/rohanraj7316/hose/api/resources/health"
	"github.com/rohanraj7316/hose/configs"
	"github.com/rohanraj7316/middleware/libs/response"

	"github.com/gofiber/fiber/v2"
)

type Router func(fiber.Router)

type Route struct {
	path   string
	router Router
}

type RouteHandler struct {
	app     *fiber.App
	sConfig *configs.ServerConfigStruct
}

func NewRouteHandler(app *fiber.App, sConfig *configs.ServerConfigStruct) (*RouteHandler, error) {
	return &RouteHandler{
		app:     app,
		sConfig: sConfig,
	}, nil
}

func (r *RouteHandler) NewRouter(app *fiber.App) {
	// list down all the routes and their handlers
	routes := []Route{
		{
			path:   "/health",
			router: health.NewRouter,
		},
		{
			path:   "/",
			router: cr.NewRouter,
		},
	}

	for i := 0; i < len(routes); i++ {
		routes[i].router(app.Group(routes[i].path))
	}

	app.Use("*", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("Cannot %s %s", c.Method(), c.Path()) // Cannot GET /healths
		return response.NewBody(c, http.StatusBadRequest, msg, nil, nil)
	})
}
