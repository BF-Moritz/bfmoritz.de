package main

import (
	"context"
	"fmt"
	"server/router"
	"server/types"
	"server/vars"
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"gopkg.in/alecthomas/kingpin.v2"

	logLib "github.com/BF-Moritz/log.lib.go"
	"github.com/BF-Moritz/log.lib.go/enum"
	logMiddleware "github.com/BF-Moritz/log.lib.go/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

var (
	config types.ConfigType

	verboseLevel = kingpin.Flag("verbose", "the level of verbosity (0-3)").Default("1").Short('v').Uint32()

	configFile = kingpin.Flag("config", "Path to configfile.").Required().File()
)

func main() {

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

	// --- init YouTube

	vars.YouTube, err = youtube.NewService(context.Background(), option.WithAPIKey(config.YouTube.ApiKey))
	if err != nil {
		vars.Logger.LogFatal("main", "failed to create youtube service")
		return
	}

	// --- init sql

	// TODO
	vars.Conn, err = sqlx.Connect(
		vars.Config.Sql.Driver,
		vars.Config.Sql.User+":"+vars.Config.Sql.Password+"@tcp("+vars.Config.Sql.Host+":"+strconv.Itoa(vars.Config.Sql.Port)+")/"+vars.Config.Sql.Name+"?charset=utf8mb4&parseTime=true&columnsWithAlias=true",
	)

	// --- init echo

	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	e.Use(logMiddleware.MakeEchoMiddleware(vars.Logger))

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: config.Server.Origins,
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	e.Pre(middleware.AddTrailingSlash())

	router.MakeRoutes(e)

	vars.Logger.LogInfo("main", "start listening on port %d", config.Server.Port)
	err = e.Start(fmt.Sprintf(":%d", config.Server.Port))
	if err != nil {
		vars.Logger.LogFatal("main", "failed to listen on Port %d (%s)", config.Server.Port, err)
	}
}
