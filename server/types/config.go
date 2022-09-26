package types

type ConfigType struct {
	Server ServerConfigType
}

type ServerConfigType struct {
	Port    int     `mapstructure:"port"`
	WebRoot *string `mapstructure:"webroot"`
}
