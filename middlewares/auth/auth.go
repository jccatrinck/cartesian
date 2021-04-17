package auth

import (
	"github.com/jccatrinck/cartesian/libs/env"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var keyAuthConfig = middleware.KeyAuthConfig{
	KeyLookup:  "header:" + echo.HeaderAuthorization,
	AuthScheme: "Bearer",
	Skipper: func(echo.Context) bool {
		apiKey := env.Get("API_KEY", "")
		return apiKey == ""
	},
	Validator: func(key string, c echo.Context) (bool, error) {
		apiKey := env.Get("API_KEY", "")
		return key == apiKey, nil
	},
}

// Middleware instance for authentication based on API_KEY env var
var Middleware = middleware.KeyAuthWithConfig(keyAuthConfig)
