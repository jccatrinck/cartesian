package main

import (
	"log"
	"net/http"

	"github.com/jccatrinck/cartesian/handlers"
	"github.com/jccatrinck/cartesian/libs/env"
	"github.com/jccatrinck/cartesian/libs/redis"
	"github.com/jccatrinck/cartesian/middlewares/auth"
	"github.com/jccatrinck/cartesian/services"
	"github.com/jccatrinck/cartesian/storage"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	// Version has its value set in build time: -ldflags "-X main.Version=$build_date"
	// Do not change the initializer to a function call or refers to other variables.
	Version string
)

func main() {
	// Configure dependencies
	configure()

	e := echo.New()

	// Add middlewares
	e.Use(middleware.RequestID())
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	e.GET("/", func(c echo.Context) (err error) {
		return c.JSON(http.StatusOK, Version)
	})

	// Create api subroute
	api := e.Group("/api", auth.Middleware)
	handlers.Create(api)

	// Starts the API
	port := env.Get("API_PORT", "9000")
	err := e.Start(":" + port)

	// Log API errors
	e.Logger.Fatal(err)
}

func configure() {
	err := redis.Configure()

	if err != nil {
		log.Fatal(err)
	}

	err = storage.Configure()

	if err != nil {
		log.Fatal(err)
	}

	err = services.Configure()

	if err != nil {
		log.Fatal(err)
	}
}
