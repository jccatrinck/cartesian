package cache

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/jccatrinck/cartesian/libs/redis"

	"github.com/labstack/echo/v4"
	_ "github.com/pingcap/parser/test_driver"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	err := redis.Configure()

	if err != nil {
		log.Fatal(err)
	}

	os.Exit(m.Run())
}

func TestCache(t *testing.T) {
	handler, c := newHander()
	err := handler(c)
	assert.NoError(t, err)
}

func TestCached(t *testing.T) {
	handler, c := newHander()
	err := handler(c)
	assert.NoError(t, err)
}

func newHander() (handler echo.HandlerFunc, c echo.Context) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c = e.NewContext(req, rec)
	handler = Cache(func(c echo.Context) error {
		return c.String(http.StatusOK, "test")
	})

	return
}
