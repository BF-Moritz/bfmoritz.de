package types

import dbLib "github.com/BF-Moritz/db.lib.go"

type ConfigType struct {
	Sql dbLib.ConfigType `mapstructure:"sql"`
}
