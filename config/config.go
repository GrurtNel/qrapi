package config

import (
	"qrapi/x/logger"

	"github.com/BurntSushi/toml"
)

type Config struct {
	ServerConfig *serverConfig `toml:"server"`
	LogConfig    *logConfig    `toml:"log"`
	StaticConfig staticConfig  `toml:"static"`
	DBConfig     dbConfig      `toml:"db"`
	PushConfig   pushConfig    `toml:"push"`
}

var conf *Config
var configLog = logger.NewLogger("config")

func init() {
	if _, err := toml.DecodeFile("./qrapi-dev.toml", &conf); err != nil {
		// handle error
		configLog.Errorf("read config error %s", err)
	}
	ServerConfig = conf.ServerConfig
	LogConfig = conf.LogConfig
	StaticConfig = conf.StaticConfig
	DBConfig = conf.DBConfig
	PushConfig = conf.PushConfig
	{
		PushConfig.initPush()
		LogConfig.initLog()
		StaticConfig.initStatic()
		DBConfig.initDB()
	}
}
