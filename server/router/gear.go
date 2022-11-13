package router

import (
	"server/handler/gear"

	"github.com/labstack/echo/v4"
)

func makeGearRoutes(g *echo.Group) {
	g.GET("/", gear.NewGetGearHandler())
}
