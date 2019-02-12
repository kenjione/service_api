package config

import (
	"encoding/json"
	"io/ioutil"
)

type ServerConfig struct {
	ServerPort       string
	DatabaseHost     string
	DatabaseUser     string
	DatabasePassword string
}

func MakeServerConfigFromFile(jsonFile string) (*ServerConfig, error) {
	b, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		return nil, err
	}
	return MakeServerConfig(b)
}

func MakeServerConfig(b []byte) (*ServerConfig, error) {
	c := &ServerConfig{}
	err := json.Unmarshal(b, c)

	return c, err
}
