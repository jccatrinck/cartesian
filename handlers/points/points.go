package points

import (
	"net/http"

	"github.com/jccatrinck/cartesian/services/points"
	"github.com/jccatrinck/cartesian/services/points/model"

	"github.com/labstack/echo/v4"
)

// Request for Handler
type Request struct {
	X        int  `query:"x"`
	Y        int  `query:"y"`
	Distance int  `query:"distance"`
	Pretty   bool `query:"pretty"`
}

// Handler returns points by distance of a point as JSON array
func Handler(c echo.Context) (err error) {
	ctx := c.Request().Context()

	request := Request{}

	err = c.Bind(&request)

	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, nil)
	}

	point := model.Point{
		X: request.X,
		Y: request.Y,
	}

	distance := request.Distance

	relativePoints, err := points.GetPointsByDistance(ctx, point, distance)

	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, nil)
	}

	if len(relativePoints) == 0 {
		return c.JSON(http.StatusNoContent, nil)
	}

	if request.Pretty {
		return c.JSONPretty(http.StatusOK, relativePoints, "\t")
	}

	return c.JSON(http.StatusOK, relativePoints)
}
