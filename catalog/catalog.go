package catalog

import (
	"encoding/json"
	"fmt"
	"github.com/colesnodgrass/abcdk/protocol"
	"os"
)

func FromFile(path string) (protocol.ConfiguredAirbyteCatalog, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return protocol.ConfiguredAirbyteCatalog{}, fmt.Errorf("failed to read file: %w", err)
	}

	var catalog protocol.ConfiguredAirbyteCatalog
	if err = json.Unmarshal(data, &catalog); err != nil {
		return protocol.ConfiguredAirbyteCatalog{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return catalog, nil
}
