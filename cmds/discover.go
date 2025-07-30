package cmds

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/colesnodgrass/abcdk/config"
	"github.com/colesnodgrass/abcdk/protocol"
	"io"
)

type DiscoverCmd struct {
	Config string `type:"existingfile"`
}

func (dc *DiscoverCmd) Run(ctx context.Context, w io.Writer) error {
	cfg, err := config.FromFile(dc.Config)
	if err != nil {
		return fmt.Errorf("failed to load config %s: %w", dc.Config, err)
	}

	data, err := json.Marshal(dc.msg(cfg))
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}

	if _, err := w.Write(data); err != nil {
		return fmt.Errorf("failed to write data: %w", err)
	}

	return nil
}

func (dc *DiscoverCmd) msg(cfg config.Config) protocol.AirbyteMessage {
	streams := []protocol.AirbyteStream{
		{
			DefaultCursorField:      cfg.Cursor(),
			IsFileBased:             nil,
			IsResumable:             nil,
			JsonSchema:              cfg.Catalog(),
			Name:                    "stream",
			Namespace:               nil,
			SourceDefinedCursor:     nil,
			SourceDefinedPrimaryKey: nil,
			SupportedSyncModes:      []protocol.SyncMode{protocol.SyncModeFullRefresh},
			AdditionalProperties:    nil,
		},
	}

	catalog := protocol.AirbyteCatalog{
		Streams:              streams,
		AdditionalProperties: nil,
	}

	return protocol.AirbyteMessage{
		Catalog: &catalog,
		Type:    protocol.AirbyteMessageTypeCATALOG,
	}
}
