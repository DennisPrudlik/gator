package config

import (
	"bytes"
	"encoding/json"
	"os"
	"path"
)

type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {
	homePath, err := os.UserHomeDir()
	if err != nil {
		return Config{}, err
	}
	configPath := path.Join(homePath, ".gatorconfig.json")
	data, err := os.ReadFile(configPath)
	if err != nil {
		return Config{}, err
	}
	var config Config
	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&config)
	if err != nil {
		return Config{}, err
	}
	return config, nil

}

func (c *Config) SetUser(userName string) error {
	homePath, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	c.CurrentUserName = userName
	configPath := path.Join(homePath, ".gatorconfig.json")
	data, err := json.Marshal(c)
	if err != nil {
		return err
	}
	err = os.WriteFile(configPath, data, 0644)
	if err != nil {
		return err
	}
	return nil
}
