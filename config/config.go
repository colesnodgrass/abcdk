package config

import (
	"encoding/json"
	"fmt"
	"github.com/colesnodgrass/abcdk/dataset"
	"os"
	"strings"
)

type Config struct {
	Check              string `json:"check"`
	Discover           string `json:"discover"`
	Read               string `json:"read"`
	Write              string `json:"write"`
	Data               string `json:"data"`
	OverrideCursor     string `json:"override-cursor"`
	OverrideRequired   string `json:"override-required"`
	OverrideProperties string `json:"override-properties"`
	OverrideRecords    string `json:"override-records"`
	dataSet            dataset.DataSet
}

func (c *Config) Cursor() []string {
	return c.dataSet.Catalog.Cursor
}

func (c *Config) Catalog() map[string]any {
	return map[string]any{
		"type":       "object",
		"$schema":    "http://json-schema.org/draft-07/schema#",
		"required":   c.dataSet.Catalog.Required,
		"properties": c.dataSet.Catalog.Properties,
	}
}

func (c *Config) Records() dataset.Records {
	return c.dataSet.Records
}

func FromFile(path string) (Config, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return Config{}, fmt.Errorf("failed to read file %s: %w", path, err)
	}

	var cfg Config
	if err := json.Unmarshal(content, &cfg); err != nil {
		return Config{}, fmt.Errorf("failed to unmarshal config %s: %w", path, err)
	}

	switch cfg.Data {
	case dataset.Movies.Name:
		cfg.dataSet = dataset.Movies
	case dataset.Moviez.Name:
		cfg.dataSet = dataset.Moviez
	case dataset.Games.Name:
		cfg.dataSet = dataset.Games
	case dataset.Sprinters.Name:
		cfg.dataSet = dataset.Sprinters
	case dataset.Advanced.Name:
		{
			var cursor string
			if err := json.Unmarshal([]byte(cfg.OverrideCursor), &cursor); err != nil {
				return Config{}, fmt.Errorf("failed to unmarshal override cursor: %w", err)
			}
			dataset.Advanced.Catalog.Cursor = strings.Split(cursor, ",")
		}

		{
			var required string
			if err := json.Unmarshal([]byte(cfg.OverrideRequired), &required); err != nil {
				return Config{}, fmt.Errorf("failed to unmarshal override required: %w", err)
			}
			dataset.Advanced.Catalog.Required = strings.Split(required, ",")
		}

		{
			var props map[string]any
			if err := json.Unmarshal([]byte(cfg.OverrideProperties), &props); err != nil {
				return Config{}, fmt.Errorf("failed to unmarshal override properties: %w", err)
			}
			dataset.Advanced.Catalog.Properties = props
		}

		{
			var records dataset.Records
			if err := json.Unmarshal([]byte(cfg.OverrideRecords), &records); err != nil {
				return Config{}, fmt.Errorf("failed to unmarshal override records: %w", err)
			}
			dataset.Advanced.Records = records
		}

		cfg.dataSet = dataset.Advanced
	default:
		return Config{}, fmt.Errorf("unknown data type %s", cfg.Data)
	}

	return cfg, nil
}
