package config

import (
	"encoding/json"
	"net"
	"os"
)

// HTTP ...
type HTTP struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

// GRPC ...
type GRPC struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

// DB ...
type DB struct {
	DSN    string `json:"dsn"`
	Source string `json:"source"`
}

// Config ...
type Config struct {
	HTTP      HTTP   `json:"http"`
	GRPC      GRPC   `json:"grpc"`
	DB        DB     `json:"db"`
	UrlLength int    `json:"url_length"`
	UrlPrefix string `json:"url_prefix"`
}

// NewConfig ...
func NewConfig(path string) (*Config, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	config := &Config{}
	err = json.Unmarshal(file, &config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func (g *GRPC) GetAddress() string {
	return net.JoinHostPort(g.Host, g.Port)
}

func (h *HTTP) GetAddress() string {
	return net.JoinHostPort(h.Host, h.Port)
}
