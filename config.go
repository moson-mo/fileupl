package main

import (
	"os"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Url           string `toml:"url"`
	ListenAddress string `toml:"listen_address"`
	APIKey        string `toml:"api_key"`
	Directory     string `toml:"directory"`
	MaxMB         int    `toml:"max_mb"`
}

func loadConfig(path string) (Config, error) {
	conf := Config{
		Url:           "http://127.0.0.1:9999",
		Directory:     "/tmp/fileupl",
		ListenAddress: ":9999",
		MaxMB:         10,
	}

	b, err := os.ReadFile(path)
	if err != nil {
		return conf, err
	}

	err = toml.Unmarshal(b, &conf)
	if err != nil {
		return conf, nil
	}

	return conf, nil
}
