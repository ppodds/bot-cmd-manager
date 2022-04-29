package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Config struct {
	Token string `json:"token"`
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) Load() bool {
	configFile, err := os.Open("config.json")
	if os.IsNotExist(err) {
		return false
	}
	bytesContent, _ := ioutil.ReadAll(configFile)
	json.Unmarshal(bytesContent, c)
	return true
}

func (c *Config) Save() {
	bytesContent, _ := json.Marshal(c)
	ioutil.WriteFile("config.json", bytesContent, 0644)
}