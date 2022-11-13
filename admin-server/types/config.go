package types

type ConfigType struct {
	Server         ServerConfigType  `mapstructure:"server"`
	VideoCacheTime uint32            `mapstructure:"videoCacheTime"`
	YouTube        YouTubeConfigType `mapstructure:"youtube"`
	Sql            SqlConfigType     `mapstructure:"sql"`
}

type ServerConfigType struct {
	Port    int      `mapstructure:"port"`
	WebRoot *string  `mapstructure:"webroot"`
	Origins []string `mapstructure:"origins"`
}

type YouTubeConfigType struct {
	ApiKey           string `mapstructure:"apikey"`
	UploadPlaylistID string `mapstructure:"uploadPlaylistID"`
}

type SqlConfigType struct {
	Driver   string `mapstructure:"driver"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Name     string `mapstructure:"name"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}
