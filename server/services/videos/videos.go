package videos

import (
	"fmt"
	"server/daos/youtube"
	"server/types"
	"server/vars"
	"time"
)

type Service struct {
	LastUpdated *time.Time
	Videos      []types.VideoType
	youtubeDAO  youtube.DAOInterface
}

type ServiceInterface interface {
	GetVideos() (videos []types.VideoType, timeStamp time.Time, err error)
}

func NewService() ServiceInterface {
	return &Service{
		Videos:      []types.VideoType{},
		LastUpdated: nil,
		youtubeDAO:  youtube.NewDAO(),
	}
}

func (s *Service) GetVideos() (videos []types.VideoType, timeStamp time.Time, err error) {

	// cache is not older than in config specified
	if s.LastUpdated != nil && time.Now().UTC().Sub(*s.LastUpdated) < time.Second*time.Duration(vars.Config.VideoCacheTime) {
		return s.Videos, *s.LastUpdated, nil
	}

	vars.Logger.LogInfo("service:videos.GetVideos()", "updating video cache")

	// get videos from youtube
	rawVideos, err := s.youtubeDAO.GetVideos()
	if err != nil {
		vars.Logger.LogError("service:videos.GetVideos()", "failed to get youtube videos (%s)", err)
		return nil, time.Time{}, err
	}

	// get tags
	tagsMap, err := s.youtubeDAO.GetVideosTags()
	if err != nil {
		vars.Logger.LogError("service:videos.GetVideos()", "failed to get youtube tags (%s)", err)
		return nil, time.Time{}, err
	}

	// get meta
	metaMap, err := s.youtubeDAO.GetVideosMeta()
	if err != nil {
		vars.Logger.LogError("service:videos.GetVideos()", "failed to get youtube meta (%s)", err)
		return nil, time.Time{}, err
	}

	videos = make([]types.VideoType, 0, len(rawVideos))

	for _, rawVideo := range rawVideos {
		if rawVideo.Status.PrivacyStatus != "public" {
			continue
		}

		publishedAt, err := time.Parse(time.RFC3339, rawVideo.Snippet.PublishedAt)
		if err != nil {
			vars.Logger.LogError("service:videos.GetVideos()", "failed to parse time (%s)", err)
			return nil, time.Time{}, err
		}

		tags := []string{"video", "livestream"}
		if t, ok := tagsMap[rawVideo.ContentDetails.VideoId]; ok {
			tags = t
		}

		// var meta *types.VideoMetaType = nil
		var desc types.LanguageStringType = types.LanguageStringType{
			German:  "Hallo, das ist eine liste:\n1. punkt\n2. strich",
			English: "Hello, this is a list:\n1. point\n2. line",
		}

		var meta *types.VideoMetaType = &types.VideoMetaType{
			Description: &desc,
			Correction:  &desc,
		}
		if m, ok := metaMap[rawVideo.ContentDetails.VideoId]; ok {
			meta = &m
		}

		videos = append(videos, types.VideoType{
			ID:           rawVideo.ContentDetails.VideoId,
			URL:          fmt.Sprintf("https://youtube.com/watch?v=%s", rawVideo.ContentDetails.VideoId),
			Title:        rawVideo.Snippet.Title,
			ChannelID:    rawVideo.Snippet.ChannelId,
			ChannelTitle: rawVideo.Snippet.ChannelTitle,
			Description:  rawVideo.Snippet.Description,
			PlaylistID:   rawVideo.Snippet.PlaylistId,
			PublishedAt:  publishedAt,
			Thumbnail: types.ThumbnailType{
				URL:    rawVideo.Snippet.Thumbnails.Maxres.Url,
				Width:  uint32(rawVideo.Snippet.Thumbnails.Maxres.Width),
				Height: uint32(rawVideo.Snippet.Thumbnails.Maxres.Height),
			},
			Tags: tags,
			Meta: meta,
		})
	}

	s.Videos = videos
	timeStamp = time.Now().UTC()
	s.LastUpdated = &timeStamp

	return
}
