package points

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/jccatrinck/cartesian/services/points"
	"github.com/jccatrinck/cartesian/services/points/model"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/pingcap/parser/test_driver"
	"github.com/stretchr/testify/assert"
)

func init() {
	middleware.DefaultLoggerConfig.Skipper = func(echo.Context) bool {
		return true
	}
}

func TestHandler(t *testing.T) {
	route := "/points?distance=0&x=0&y=0"

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, route, strings.NewReader(""))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	storage := MockStorage{}

	file, err := ioutil.TempFile("", "dummy")
	assert.NoError(t, err)

	err = os.Setenv("POINTS_FILE", file.Name())
	assert.NoError(t, err)

	err = points.Configure(storage)
	assert.NoError(t, err)

	err = Handler(c)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)

	body := rec.Body.Bytes()

	points := []model.RelativePoint{}

	err = json.Unmarshal(body, &points)
	assert.NoError(t, err)
	assert.NotEmpty(t, points)
}

func TestHandlerInvalid(t *testing.T) {
	routeInvalid := "/points?distance=A&x=0.0&y=0.0"

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, routeInvalid, strings.NewReader(""))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := Handler(c)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.NotEmpty(t, rec.Body.String())
}

func TestRequest(t *testing.T) {
	route := "/points?distance=9&x=9&y=9&pretty=true"

	request := Request{}

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, route, nil)
	rec := httptest.NewRecorder()
	context := e.NewContext(req, rec)

	binder := &echo.DefaultBinder{}

	err := binder.BindQueryParams(context, &request)
	assert.NoError(t, err)

	assert.EqualValues(t, Request{9, 9, 9, true}, request)
}

type MockStorage struct{}

func (s MockStorage) LoadPoints(reader io.ReadSeeker) error {
	return nil
}

func (s MockStorage) GetPointsByDistance(point model.Point, distance int) ([]model.RelativePoint, error) {
	return []model.RelativePoint{
		{
			Point: model.Point{
				X: 0,
				Y: 0,
			},
			Distance: 0,
		},
	}, nil
}
