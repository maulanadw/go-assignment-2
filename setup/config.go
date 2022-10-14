package setup

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
}

const FILENAME = "setup/config.json"

func GetConfig() *Config {
	configFile, err := ioutil.ReadFile(FILENAME)
	if err != nil {
		msg := fmt.Sprintf("Failed reading config file: %s", err)
		panic(msg)
	}

	var config Config
	err = json.Unmarshal(configFile, &config)
	if err != nil {
		msg := fmt.Sprintf("Failed parsing config file: %s", err)
		panic(msg)
	}

	return &config
}
