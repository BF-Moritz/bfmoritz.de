package gear

import (
	"admin-server/services/gear"

	"github.com/labstack/echo/v4"
)

func NewPatchGearHandler() echo.HandlerFunc {
	return newPatchGearHandlerWithDeps(gear.NewService())
}

func newPatchGearHandlerWithDeps(gearService gear.ServiceInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		// TODO
		return c.JSON(200, nil)
	}
}
