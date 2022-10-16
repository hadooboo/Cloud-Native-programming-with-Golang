package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

var (
	DefaultDatabaseType      = "mongodb"
	DefaultDatabaseConn      = "mongodb://127.0.0.1"
	DefaultApiserverEndpoint = "localhost:8181"
)

type Config struct {
	Database  DatabaseConfig  `yaml:"database"`
	Apiserver ApiserverConfig `yaml:"apiserver"`
}

type DatabaseConfig struct {
	Type string `yaml:"type"`
	Conn string `yaml:"conn"`
}

type ApiserverConfig struct {
	Endpoint string `yaml:"endpoint"`
}

func NewConfig(filepath string) (*Config, error) {
	config := &Config{
		Database: DatabaseConfig{
			Type: DefaultDatabaseType,
			Conn: DefaultDatabaseConn,
		},
		Apiserver: ApiserverConfig{
			Endpoint: DefaultApiserverEndpoint,
		},
	}

	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	if err := yaml.NewDecoder(file).Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}
