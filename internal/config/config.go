package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {
	var config Config
	p, err := getConfigFilePath()
	if err != nil {
		return config, err
	}

	cfgFile, err := os.Open(p)
	if err != nil {
		return config, err
	}

	defer cfgFile.Close()

	if err := json.NewDecoder(cfgFile).Decode(&config); err != nil {
		return config, err
	}

	return config, nil
}

func (c *Config) SetUser(current_user string) error {
	c.CurrentUserName = current_user
	if err := write(*c); err != nil {
		return err
	}

	return nil
}

func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	p := filepath.Join(home, configFileName)
	return p, nil
}

func write(cfg Config) error {
	p, err := getConfigFilePath()
	if err != nil {
		return fmt.Errorf("error getting config path: %w", err)
	}

	cfgFile, err := os.Create(p)
	if err != nil {
		return fmt.Errorf("error opening file: %w", err)
	}

	defer cfgFile.Close()

	if err := json.NewEncoder(cfgFile).Encode(cfg); err != nil {
		return fmt.Errorf("error encoding json: %w", err)
	}
	return nil
}
