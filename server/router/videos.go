package router

import (
	"server/handler/videos"

	"github.com/labstack/echo/v4"
)

func makeVideoRoutes(g *echo.Group) {
	g.GET("/", videos.NewGetVideosHandler())
}
