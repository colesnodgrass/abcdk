package config

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/colesnodgrass/abcdk/dataset"
)

type Config struct {
	Check              string `json:"check"`
	Discover           string `json:"discover"`
	Read               string `json:"read"`
	Write              string `json:"write"`
	Data               string `json:"data"`
	CustomCursor     string `json:"custom_cursor"`
	CustomRequired   string `json:"custom_required"`
	CustomProperties string `json:"custom_properties"`
	CustomRecords    string `json:"custom_records"`
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

func (c *Config) Stream() string {
	return "stream-" + strings.ToLower(c.Data)
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
	case dataset.MoviesMapperFilter.Name:
		cfg.dataSet = dataset.MoviesMapperFilter
	case dataset.MoviesFilmOnly.Name:
		cfg.dataSet = dataset.MoviesFilmOnly
	case dataset.Games.Name:
		cfg.dataSet = dataset.Games
	case dataset.Sprinters.Name:
		cfg.dataSet = dataset.Sprinters
	case dataset.Custom.Name:
		{
			var cursor string
			if err := json.Unmarshal([]byte(cfg.CustomCursor), &cursor); err != nil {
				return Config{}, fmt.Errorf("failed to unmarshal custom cursor: %w", err)
			}
			dataset.Custom.Catalog.Cursor = strings.Split(cursor, ",")
		}

		{
			var required string
			if err := json.Unmarshal([]byte(cfg.CustomRequired), &required); err != nil {
				return Config{}, fmt.Errorf("failed to unmarshal custom required: %w", err)
			}
			dataset.Custom.Catalog.Required = strings.Split(required, ",")
		}

		{
			var props map[string]any
			if err := json.Unmarshal([]byte(cfg.CustomProperties), &props); err != nil {
				return Config{}, fmt.Errorf("failed to unmarshal custom properties: %w", err)
			}
			dataset.Custom.Catalog.Properties = props
		}

		{
			var records dataset.Records
			if err := json.Unmarshal([]byte(cfg.CustomRecords), &records); err != nil {
				return Config{}, fmt.Errorf("failed to unmarshal custom records: %w", err)
			}
			dataset.Custom.Records = records
		}

		cfg.dataSet = dataset.Custom
	default:
		return Config{}, fmt.Errorf("unknown data type %s", cfg.Data)
	}

	return cfg, nil
}
