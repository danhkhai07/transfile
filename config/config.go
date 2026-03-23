package config

type ServerConfig struct {
	Port string
}

type ClientConfig struct {
	Port string
	ServerAddr string
}
