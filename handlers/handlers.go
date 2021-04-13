package handlers

import (
	"github.com/jccatrinck/cartesian/handlers/points"

	"github.com/labstack/echo/v4"
)

// Create all API handlers
func Create(g *echo.Group) {
	g.GET("/points", points.Handler)
}