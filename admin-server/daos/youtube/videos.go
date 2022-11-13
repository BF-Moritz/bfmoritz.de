package youtube

import (
	"admin-server/types"
	"admin-server/vars"

	"google.golang.org/api/googleapi"
	"google.golang.org/api/youtube/v3"
)

type DAO struct {
}

type DAOInterface interface {
	GetVideos() (videos []youtube.PlaylistItem, err error)
	GetVideosTags() (tags map[string][]string, err error)
	GetVideosMeta() (meta map[string]types.VideoMetaType, err error)
}

func NewDAO() DAOInterface {
	return &DAO{}
}

func (d *DAO) GetVideos() (videos []youtube.PlaylistItem, err error) {
	call := vars.YouTube.PlaylistItems.List([]string{"contentDetails", "id", "snippet", "status"})

	query := 0

	resp, err := call.Do(googleapi.QueryParameter("maxResults", "10"), googleapi.QueryParameter("playlistId", vars.Config.YouTube.UploadPlaylistID))
	if err != nil {
		vars.Logger.LogError("dao:youtube.GetVideos()", "failed to query videos <%d> (%s)", query, err)
		return nil, err
	}

	videos = make([]youtube.PlaylistItem, 0, resp.PageInfo.TotalResults)

	for _, item := range resp.Items {
		videos = append(videos, *item)
	}

	nextToken := resp.NextPageToken

	for nextToken != "" {
		query++

		resp, err := call.Do(
			googleapi.QueryParameter("maxResults", "10"),
			googleapi.QueryParameter("playlistId", vars.Config.YouTube.UploadPlaylistID),
			googleapi.QueryParameter("pageToken", nextToken),
		)
		if err != nil {
			vars.Logger.LogError("dao:youtube.GetVideos()", "failed to query videos <%d> (%s)", query, err)
			return nil, err
		}

		nextToken = resp.NextPageToken

		for _, item := range resp.Items {
			videos = append(videos, *item)
		}
	}

	return
}

func (d *DAO) GetVideosTags() (tags map[string][]string, err error) {
	// TODO make db request
	return make(map[string][]string), nil
}

func (d *DAO) GetVideosMeta() (meta map[string]types.VideoMetaType, err error) {
	// TODO make db request
	return make(map[string]types.VideoMetaType), nil
}
