package gear

import (
	"admin-server/services/gear"

	"github.com/labstack/echo/v4"
)

func NewPostGearHandler() echo.HandlerFunc {
	return newPostGearHandlerWithDeps(gear.NewService())
}

func newPostGearHandlerWithDeps(gearService gear.ServiceInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		// TODO
		return c.JSON(200, nil)
	}
}
