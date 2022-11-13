package router

import (
	"github.com/labstack/echo/v4"
)

func MakeRoutes(e *echo.Echo) {
	g := e.Group("/api/v1")

	makePingRoutes(g.Group("/ping"))
	makeVideoRoutes(g.Group("/videos"))
	makeGearRoutes(g.Group("/gear"))
}
