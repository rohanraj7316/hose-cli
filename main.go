package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/helmet/v2"
	"github.com/rohanraj7316/hose/api/routes"
	"github.com/rohanraj7316/hose/configs"
	"github.com/rohanraj7316/logger"
)

func main() {

	os.Setenv("MODULE_NAME", "crypto")
	os.Setenv("PRODUCT_NAME", "hose")
	os.Setenv("WAIT_TIME_BEFORE_KILL", "60s")
	os.Setenv("PORT", "8080")
	os.Setenv("ALLOW_CORS_METHODS", "GET,POST")
	os.Setenv("ALLOW_CORS_ORIGIN", "*")

	// can be used to terminate the server using done
	pCtx := context.Background()
	ctx, cancel := context.WithCancel(pCtx)
	defer cancel()

	// initialize logger
	err := logger.Configure()
	if err != nil {
		log.Panic(err)
	}

	config, err := configs.NewServerConfig()
	if err != nil {
		logger.Error(err.Error())
		cancel()
	}

	app := fiber.New(config.ServerConfig)

	// adding middleware
	app.Use(cors.New(config.CorsConfig))
	app.Use(helmet.New())

	// initialize router
	r, err := routes.NewRouteHandler(app, config)
	if err != nil {
		logger.Error(err.Error())
		cancel()
	}

	r.NewRouter(app)

	logger.Info("successful starting server :)")

	err = app.Listen(fmt.Sprintf(":%s", config.Port))
	if err != nil {
		logger.Error(err.Error())
		cancel()
	}

	cChannel := make(chan os.Signal, 2)
	signal.Notify(cChannel, os.Interrupt, syscall.SIGTERM)

bLoop:
	for {
		select {
		case <-ctx.Done():
			break bLoop

		case <-cChannel:
			logger.Warn("catch interrupted signal")
			time.Sleep(config.WaitTimeBeforeKill)
			break bLoop
		}
	}

	err = app.Shutdown()
	if err != nil {
		logger.Error(err.Error())
	}

	logger.Warn("shutting down the server :(")
}
