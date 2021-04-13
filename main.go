package main

import (
	"github.com/jccatrinck/cartesian/libs/env"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.Logger.Fatal(e.Start(":" + env.Get("API_PORT", "9000")))
}
