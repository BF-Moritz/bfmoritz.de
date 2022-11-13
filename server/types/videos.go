package types

import "time"

type VideoType struct {
	ID           string         `json:"id"`
	URL          string         `json:"url"`
	Title        string         `json:"title"`
	Description  string         `json:"description"`
	ChannelID    string         `json:"channel_id"`
	ChannelTitle string         `json:"channel_title"`
	PlaylistID   string         `json:"playlist_id"`
	PublishedAt  time.Time      `json:"published_at"`
	Thumbnail    ThumbnailType  `json:"thumbnail"`
	Tags         []string       `json:"tags"`
	Meta         *VideoMetaType `json:"meta"`
}

type ThumbnailType struct {
	URL    string `json:"url"`
	Width  uint32 `json:"width"`
	Height uint32 `json:"height"`
}

type VideoMetaType struct {
	Description  *LanguageStringType  `json:"description"`
	Correction   *LanguageStringType  `json:"correction"`
	TextSources  []LanguageStringType `json:"text_sources"`
	VideoSources []LanguageStringType `json:"video_sources"`
	AudioSources []LanguageStringType `json:"audio_sources"`
}

type LanguageStringType struct {
	German  string `json:"de-DE"`
	English string `json:"en-US"`
}
