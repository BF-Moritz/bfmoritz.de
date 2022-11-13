package ping

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func NewGetPingHandler() echo.HandlerFunc {
	return newGetPingHandlerWithDeps()
}

func newGetPingHandlerWithDeps() echo.HandlerFunc {
	return func(c echo.Context) error {

		return c.NoContent(http.StatusOK)
	}
}
