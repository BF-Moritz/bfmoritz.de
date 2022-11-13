package videos

import (
	"admin-server/services/videos"
	"admin-server/types"
	"admin-server/vars"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func NewGetVideosHandler() echo.HandlerFunc {
	return newGetVideosHandlerWithDeps(videos.NewService())
}

func newGetVideosHandlerWithDeps(videoSevice videos.ServiceInterface) echo.HandlerFunc {
	return func(c echo.Context) error {

		videos, ts, err := videoSevice.GetVideos()
		if err != nil {
			vars.Logger.LogError("handler:videos.newGetVideosHandlerWithDeps()", "failed to get videos (%s)", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		return c.JSON(200, outType{
			TimeStamp: ts,
			Videos:    videos,
		})
	}
}

type outType struct {
	TimeStamp time.Time         `json:"timestamp"`
	Videos    []types.VideoType `json:"videos"`
}
