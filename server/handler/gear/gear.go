package gear

import (
	"net/http"
	"server/services/gear"
	"server/vars"

	"github.com/labstack/echo/v4"
)

func NewGetGearHandler() echo.HandlerFunc {
	return newGetGearHandlerWithDeps(gear.NewService())
}

func newGetGearHandlerWithDeps(gearService gear.ServiceInterface) echo.HandlerFunc {
	return func(c echo.Context) error {

		gear, err := gearService.GetGear()
		if err != nil {
			vars.Logger.LogError("handler:videos.newGetGearHandlerWithDeps()", "failed to get videos (%s)", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		return c.JSON(200, gear)
	}
}
