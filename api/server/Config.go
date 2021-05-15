package server

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// Config - Represents a Configuration
type Config struct {
	Version int32

	DB struct {
		Redis struct {
			Host                  string `yaml:"host"`
			Password              string `yaml:"password"`
			DB                    int    `yaml:"db"`
			PoolSize              int    `yaml:"poolSize"`
			RedisScriptDefinition string `yaml:"redisScriptDefinition"`
			RedisScriptDB         int    `yaml:"redisScriptDB"`
		}
	}
	API struct {
		URL                       string `yaml:"url"`
		Name                      string `yaml:"name"`
		MaxConnsPerHost           int    `yaml:"maxConnsPerHost"`
		MaxIdemponentCallAttempts int    `yaml:"maxIdemponentCallAttempts"`
		PostScoreURI              string `yaml:"postscoreUri"`
		GetLeaderBoardURI         string `yaml:"getleaderboardUri"`
	}
}

// LoadConfigFromFile - Load configuration from file
func LoadConfigFromFile(filename string) *Config {
	config := &Config{}

	buffer, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	err = yaml.Unmarshal([]byte(buffer), &config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return config
}
