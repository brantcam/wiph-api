package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"wiph-api/store/postgres"
)

// Config ...
type Config struct {
	Postgres postgres.Config
}

// LoadConfigFromEnv ...
func LoadConfigFromEnv() (*Config, error) {
	buf, err := ioutil.ReadFile("./config/local.json")
	if err != nil {
		return nil, err
	}
	c := new(Config)
	if err := json.Unmarshal(buf, &c.Postgres); err != nil {
		return nil, err
	}

	if value, exists := os.LookupEnv("DB_USER"); exists {
		c.Postgres.Username = value
	}
	if value, exists := os.LookupEnv("DB_NAME"); exists {
		c.Postgres.Name = value
	}
	if value, exists := os.LookupEnv("DB_PASS"); exists {
		c.Postgres.Password = value
	}

	return c, nil
}
