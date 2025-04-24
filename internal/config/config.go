package config

import (
	"encoding/json"
	"os"
	"path"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() Config {
	p, err := getConfigFilePath()
	if err != nil {
		panic(err)
	}

	var config Config
	cfgFile, err := os.Open(p)
	if err != nil {
		panic(err)
	}

	defer cfgFile.Close()

	if err := json.NewDecoder(cfgFile).Decode(&config); err != nil {
		panic(err)
	}

	return config
}

func (c *Config) SetUser(current_user string) {
	p, err := getConfigFilePath()
	if err != nil {
		panic(err)
	}

	cfgFile, err := os.Open(p)
	if err != nil {
		panic(err)
	}

	defer cfgFile.Close()

	if err := json.NewEncoder(cfgFile).Encode(c); err != nil {
		panic(err)
	}

}

func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	p := path.Join(home, configFileName)
	return p, nil
}

func write(cfg Config) error {
	return nil
}
