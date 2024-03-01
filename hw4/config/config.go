package config

import (
	"os"
)

type Config struct {
	address  string
	filename string
}

func NewConfig() (Config, error) {
	if len(os.Args) < 3 {
		return Config{}, InvalidArguments
	}
	return Config{address: os.Args[1], filename: os.Args[2]}, nil
}

func (c *Config) GetAddress() string {
	return c.address
}

func (c *Config) GetFilename() string {
	return c.filename
}
