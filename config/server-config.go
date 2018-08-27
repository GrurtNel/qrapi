package config

type serverConfig struct {
	Host string `toml:"host"`
	Port string `toml:"port"`
}

func (s *serverConfig) String() string {
	return s.Host + ":" + s.Port
}

var ServerConfig *serverConfig
