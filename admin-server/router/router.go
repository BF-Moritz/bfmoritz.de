package router

import (
	"admin-server/middlewares"

	"github.com/labstack/echo/v4"
)

func MakeRoutes(e *echo.Echo) {
	g := e.Group("/api/v1")

	makePingRoutes(g.Group("/ping"))
	makeAuthRoutes(g.Group("/auth"))

	g.Use(middlewares.AuthMiddleware())

	makeVideoRoutes(g.Group("/videos"))
	makeGearRoutes(g.Group("/gear"))
}
