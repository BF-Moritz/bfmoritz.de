package gear

import (
	"admin-server/services/gear"

	"github.com/labstack/echo/v4"
)

func NewDeleteGearHandler() echo.HandlerFunc {
	return newDeleteGearHandlerWithDeps(gear.NewService())
}

func newDeleteGearHandlerWithDeps(gearService gear.ServiceInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		// TODO
		return c.JSON(200, nil)
	}
}
