package main

import (
	"init-db/daos/db"
	"init-db/types"
	"init-db/vars"

	dblibgo "github.com/BF-Moritz/db.lib.go"
	logLib "github.com/BF-Moritz/log.lib.go"
	"github.com/BF-Moritz/log.lib.go/enum"
	"github.com/spf13/viper"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	config types.ConfigType

	verboseLevel = kingpin.Flag("verbose", "the level of verbosity (0-3)").Default("1").Short('v').Uint32()
	configFile   = kingpin.Flag("config", "Path to configfile.").Required().File()
	forceMode    = kingpin.Flag("force", "Force a rebuild").Default("false").Short('f').Bool()
)

func main() {
	// TODO

	// --- init kingpin
	kingpin.Parse()

	if *verboseLevel > 3 || *verboseLevel < 0 {
		kingpin.FatalUsage("the level of verbosity has to be between 0 and 3")
	}

	vars.Logger = logLib.NewLogger(enum.LogLevel(*verboseLevel), true, true)

	// --- init config

	if configFile == nil {
		vars.Logger.LogFatal("main", "missing configuration file")
	}

	viper.SetConfigType("yaml")
	err := viper.ReadConfig(*configFile)
	if err != nil {
		vars.Logger.LogFatal("main", "failed to read config (%s)", err)
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		vars.Logger.LogFatal("main", "failed to unmarshal config (%s)", err)
	}

	vars.Config = config

	// --- init sql

	vars.Conn, err = dblibgo.NewConn(vars.Config.Sql, vars.Logger)
	if err != nil {
		vars.Logger.LogFatal("main", "failed to create db connection (%s)", err)
	}

	// --- init dbs

	dao := db.NewDAO()

	exists, err := dao.CheckDBs()
	if err != nil {
		vars.Logger.LogFatal("main", "failed to check dbs (%s)", err)
	}

	if exists && !*forceMode {
		vars.Logger.LogInfo("main", "databases exist, please use '-f' / '--force' to delete and rebuild them (%s)", err)
		return
	}

	if *forceMode {
		err = dao.DeleteDBs()
		if err != nil {
			vars.Logger.LogFatal("main", "failed to erase dbs (%s)", err)
		}
	}

	err = dao.CreateDBs()
	if err != nil {
		vars.Logger.LogFatal("main", "failed to create dbs (%s)", err)
	}
}
