package router

import (
	"admin-server/handler/gear"

	"github.com/labstack/echo/v4"
)

func makeGearRoutes(g *echo.Group) {
	g.GET("/", gear.NewGetGearHandler())
	g.POST("/", gear.NewPostGearHandler())
	g.PATCH("/:id", gear.NewPatchGearHandler())
	g.DELETE("/:id", gear.NewDeleteGearHandler())
}
