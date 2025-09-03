package cmds

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/colesnodgrass/abcdk/catalog"
	"github.com/colesnodgrass/abcdk/config"
	"github.com/colesnodgrass/abcdk/protocol"
)

type ReadCmd struct {
	Config  string `type:"existingfile"`
	Catalog string `type:"existingfile"`
	Stream  string `type:"existingfile" optional:""`
}

func (rc *ReadCmd) Run(ctx context.Context, w io.Writer) error {
	cfg, err := config.FromFile(rc.Config)
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	_, err = catalog.FromFile(rc.Catalog)
	if err != nil {
		return fmt.Errorf("failed to load catalog: %w", err)
	}

	records := cfg.Records()
	for _, record := range records {
		data, err := json.Marshal(rc.msgRecord(rc.record(cfg.Stream(), record)))
		if err != nil {
			return fmt.Errorf("failed to marshal record: %w", err)
		}
		if _, err := w.Write(data); err != nil {
			return fmt.Errorf("failed to write record: %w", err)
		}
	}

	data, err := json.Marshal(rc.msgState(rc.state()))
	if err != nil {
		return fmt.Errorf("failed to marshal record: %w", err)
	}
	if _, err := w.Write(data); err != nil {
		return fmt.Errorf("failed to write record: %w", err)
	}

	return nil
}

func (rc *ReadCmd) msgRecord(msg *protocol.AirbyteRecordMessage) protocol.AirbyteMessage {
	return protocol.AirbyteMessage{
		Record: msg,
		Type:   protocol.AirbyteMessageTypeRECORD,
	}
}

func (rc *ReadCmd) msgState(msg *protocol.AirbyteStateMessage) protocol.AirbyteMessage {
	return protocol.AirbyteMessage{
		State: msg,
		Type:  protocol.AirbyteMessageTypeSTATE,
	}
}

func (rc *ReadCmd) record(stream string, data map[string]any) *protocol.AirbyteRecordMessage {
	return &protocol.AirbyteRecordMessage{
		Data:      data,
		EmittedAt: 987654321000,
		Stream:    stream,
	}
}

func (rc *ReadCmd) state() *protocol.AirbyteStateMessage {
	return &protocol.AirbyteStateMessage{
		Data:             nil,
		DestinationStats: nil,
		Global:           nil,
		SourceStats:      nil,
		Stream: &protocol.AirbyteStreamState{
			StreamDescriptor:     protocol.StreamDescriptor{Name: "stream"},
			StreamState:          nil,
			AdditionalProperties: nil,
		},
		Type:                 p(protocol.AirbyteStateTypeSTREAM),
		AdditionalProperties: nil,
	}
}
