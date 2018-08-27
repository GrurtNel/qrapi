package config

import (
	"flag"
	"os"
)

type logConfig struct {
	LogDirectory string `toml:"log_dir"`
	LogToStdErr  string `toml:"log_std"`
}

var LogConfig *logConfig

func (l *logConfig) initLog() {
	if _, err := os.Stat(l.LogDirectory); err != nil {
		os.Mkdir(l.LogDirectory, os.ModeAppend)
	}
	flag.Set("alsologtostderr", l.LogToStdErr)
	flag.Set("log_dir", l.LogDirectory)
	flag.Parse()
}
