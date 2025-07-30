package cmds

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/colesnodgrass/abcdk/dataset"
	"github.com/colesnodgrass/abcdk/protocol"
	"io"
)

type SpecCmd struct{}

func (sc *SpecCmd) Run(ctx context.Context, w io.Writer) error {
	data, err := json.Marshal(sc.msg())
	if err != nil {
		return fmt.Errorf("error marshalling Spec: %v", err)
	}
	if _, err := w.Write(data); err != nil {
		return fmt.Errorf("failed to write: %w\n", err)
	}
	return nil
}

func (sc *SpecCmd) spec() *protocol.ConnectorSpecification {
	spec := map[string]any{
		"type":     "object",
		"$schema":  "http://json-schema.org/draft-07/schema#",
		"required": []string{"check", "discover", "read", "write", "data"},
		"properties": map[string]any{
			"check": map[string]any{
				"order":       0,
				"type":        "string",
				"description": "Should check pass or fail?",
				"enum":        []string{"pass", "fail"},
				"default":     "pass",
			},
			"discover": map[string]any{
				"order":       1,
				"type":        "string",
				"description": "Should discover pass or fail?",
				"enum":        []string{"pass", "fail"},
				"default":     "pass",
			},
			"read": map[string]any{
				"order":       2,
				"type":        "string",
				"description": "Should read pass or fail?",
				"enum":        []string{"pass", "fail"},
				"default":     "pass",
			},
			"write": map[string]any{
				"order":       3,
				"type":        "string",
				"description": "Should write pass or fail?",
				"enum":        []string{"pass", "fail"},
				"default":     "pass",
			},
			"data": map[string]any{
				"order":       4,
				"type":        "string",
				"description": "Which dataset should be used?",
				"enum":        []string{dataset.Movies.Name, dataset.Moviez.Name, dataset.Games.Name, dataset.Sprinters.Name, dataset.Advanced.Name},
				"default":     dataset.Movies.Name,
			},
			"override-cursor": map[string]any{
				"order":       5,
				"type":        "string",
				"description": "Cursor override (comma-delimited list). Only applicable if data is set to advanced...",
				"default":     `"id"`,
			},
			"override-required": map[string]any{
				"order":       6,
				"type":        "string",
				"description": "Required override (comma-delimited list). Only applicable if data is set to advanced...",
				"default":     `"id,name"`,
			},
			"override-properties": map[string]any{
				"order":       7,
				"type":        "string",
				"description": "Catalog override (jsonschema). Only applicable if data is set to advanced...",
				"default":     `{"id":{"type":"number"},"name":{"type":"string"}}`,
			},
			"override-records": map[string]any{
				"order":       8,
				"type":        "string",
				"description": "Record override (type TDB). Only applicable if data is set to advanced...",
				"default":     `[{"id":101,"name":"foo"},{"id":202,"name":"bar"}]`,
			},
		},
	}
	return &protocol.ConnectorSpecification{
		AdvancedAuth:                  nil,
		AuthSpecification:             nil,
		ChangelogUrl:                  nil,
		ConnectionSpecification:       spec,
		DocumentationUrl:              p("https://github.com/colesnodgrass/abcdk"),
		ProtocolVersion:               p("0.18.0"),
		SupportedDestinationSyncModes: []protocol.DestinationSyncMode{protocol.DestinationSyncModeOverwrite},
		SupportsDBT:                   false,
		SupportsIncremental:           nil,
		SupportsNormalization:         false,
		AdditionalProperties:          nil,
	}
}

func (sc *SpecCmd) msg() protocol.AirbyteMessage {
	return protocol.AirbyteMessage{
		Catalog:              nil,
		ConnectionStatus:     nil,
		Control:              nil,
		DestinationCatalog:   nil,
		Log:                  nil,
		Record:               nil,
		Spec:                 sc.spec(),
		State:                nil,
		Trace:                nil,
		Type:                 protocol.AirbyteMessageTypeSPEC,
		AdditionalProperties: nil,
	}
}

func p[T any](v T) *T {
	return &v
}
