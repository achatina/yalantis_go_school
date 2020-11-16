package config

import (
	"gopkg.in/yaml.v2"
	"os"
)

type Configuration struct {
	Server struct {
		Port string `yaml:"port"`
		Host string `yaml:"host"`
	}
	Database struct {
		Username string `yaml:"user"`
		Password string `yaml:"pass"`
		Name     string `yaml:"name"`
		Port string `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"database"`
}

func GetConfiguration() Configuration {
	f, err := os.Open("config.yml")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var cfg Configuration
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		panic(err)
	}
	return cfg
}