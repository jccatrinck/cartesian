package points

import (
	"net/http"

	"github.com/shopspring/decimal"

	"github.com/labstack/echo/v4"
)

type Request struct {
	X        decimal.Decimal `query:"x"`
	Y        decimal.Decimal `query:"y"`
	Distance decimal.Decimal `query:"distance"`
	Pretty   bool            `query:"pretty"`
}

// Handler returns all/filtered points as JSON array
func Handler(c echo.Context) (err error) {
	request := Request{}

	err = c.Bind(&request)

	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, nil)
	}

	relativePoints := []struct{}{}

	if request.Pretty {
		return c.JSONPretty(http.StatusOK, relativePoints, "\t")
	}

	return c.JSON(http.StatusOK, relativePoints)
}
