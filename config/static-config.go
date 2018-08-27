package config

import "os"

type staticConfig struct {
	UploadDirectory string `toml:"upload_dir"`
}

var StaticConfig staticConfig

func (s staticConfig) initStatic() {
	if _, err := os.Stat(s.UploadDirectory); err != nil {
		os.Mkdir(s.UploadDirectory, os.ModeAppend)
	}
}
