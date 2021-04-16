package cache

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/jccatrinck/cartesian/libs/redis"
	"github.com/labstack/echo/v4"
)

const (
	cacheDuration time.Duration = 1 * time.Minute
	cachePrefix   string        = "request-cache"
)

type cachedResponse struct {
	StatusCode int
	Body       []byte
}

// Cache middleware takes entire URL and body as key to cache JSON response
func Cache(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		if redis.Client == nil {
			return next(c)
		}

		ctx := c.Request().Context()

		url := c.Request().URL.String()

		// Request
		request := []byte{}
		if c.Request().Body != nil {
			request, err = ioutil.ReadAll(c.Request().Body)

			if err != nil {
				return c.JSON(http.StatusInternalServerError, nil)
			}
		}

		// Reset
		c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(request))

		var requestIndent string

		if len(request) > 0 {
			var prettyJSON bytes.Buffer

			err = json.Indent(&prettyJSON, request, "", " ")

			if err != nil {
				return c.JSON(http.StatusInternalServerError, nil)
			}

			requestIndent = prettyJSON.String()
		}

		key := fmt.Sprintf("%s-%s-%s", cachePrefix, url, requestIndent)

		cached, err := redis.Client.Get(ctx, key).Result()

		if err == redis.Nil {
			err = nil
		}

		if err != nil {
			return c.JSON(http.StatusInternalServerError, nil)
		}

		// Stop processing if client went away
		select {
		case <-ctx.Done():
			return nil
		default:
		}

		if cached != "" {
			response := cachedResponse{}

			err = json.Unmarshal([]byte(cached), &response)

			if err != nil {
				return
			}

			return c.JSONBlob(response.StatusCode, response.Body)
		}

		// Reponse
		response := new(bytes.Buffer)
		mw := io.MultiWriter(c.Response().Writer, response)
		writer := &bodyDumpResponseWriter{Writer: mw, ResponseWriter: c.Response().Writer}
		c.Response().Writer = writer

		err = next(c)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, nil)
		}

		cachedResp := cachedResponse{
			StatusCode: c.Response().Status,
			Body:       response.Bytes(),
		}

		cachedRespJSON, err := json.Marshal(cachedResp)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, nil)
		}

		err = redis.Client.Set(ctx, key, string(cachedRespJSON), cacheDuration).Err()

		if err != nil {
			return c.JSON(http.StatusInternalServerError, nil)
		}

		return
	}
}
