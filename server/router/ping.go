package router

import (
	"server/handler/ping"

	"github.com/labstack/echo/v4"
)

func makePingRoutes(g *echo.Group) {
	g.GET("/", ping.NewGetPingHandler())
}
