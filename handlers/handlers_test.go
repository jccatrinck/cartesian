package handlers

import (
	"testing"

	"github.com/labstack/echo/v4"
	_ "github.com/pingcap/parser/test_driver"
)

func TestCreate(t *testing.T) {
	g := echo.New().Group("/api")
	Create(g)
}
