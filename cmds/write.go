package cmds

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/colesnodgrass/abcdk/airbyte"
	"github.com/colesnodgrass/abcdk/catalog"
	"github.com/colesnodgrass/abcdk/config"
	"github.com/colesnodgrass/abcdk/dataset"
	"github.com/colesnodgrass/abcdk/protocol"
	"github.com/colesnodgrass/go-diff/diff"
)

type WriteCmd struct {
	Config  string `type:"existingfile"`
	Catalog string `type:"existingfile"`
}

func (wc *WriteCmd) Run(ctx context.Context, w io.Writer, r io.Reader) error {
	{
		if contents, err := os.ReadFile(wc.Config); err == nil {
			_ = airbyte.LogInfo(w, fmt.Sprintf("Write config (raw):\n%s\n", contents))
		}
		if contents, err := os.ReadFile(wc.Catalog); err == nil {
			_ = airbyte.LogInfo(w, fmt.Sprintf("Write catalog (raw):\n%s\n", contents))
		}
	}

	diff.Color = false

	cfg, err := config.FromFile(wc.Config)
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	_, err = catalog.FromFile(wc.Catalog)
	if err != nil {
		return fmt.Errorf("failed to load catalog: %w", err)
	}

	records := cfg.Records()
	recordsPtr := 0
	recordsLen := len(records)
	recordsFailed := 0

	decoder := json.NewDecoder(r)
	for {
		_ = airbyte.LogInfo(w, "write...")
		var msg protocol.AirbyteMessage
		if err := decoder.Decode(&msg); err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			_ = airbyte.LogError(w, fmt.Errorf("failed to decode record: %w", err))
		}
		_ = airbyte.LogInfo(w, "write 2...")

		switch msg.Type {
		case protocol.AirbyteMessageTypeRECORD:
			var expected map[string]any
			if recordsPtr < recordsLen {
				expected = records[recordsPtr]
			}
			data, _ := json.Marshal(msg.Record.Data)
			_ = airbyte.LogInfo(w, fmt.Sprintf("Write record: %s", data))

			if !checkRecord(w, expected, msg.Record.Data) {
				recordsFailed++
			}
			recordsPtr++
		case protocol.AirbyteMessageTypeSTATE:
			data, _ := json.Marshal(msg.State.Data)
			_ = airbyte.LogInfo(w, fmt.Sprintf("Write state: %s", data))
		}
	}

	for i := recordsPtr; i < recordsLen; i++ {
		if !checkRecord(w, records[i], nil) {
			recordsFailed++
		}
	}

	if recordsFailed > 0 {
		return fmt.Errorf("%d records failed", recordsFailed)
	}

	return nil
}

// checkRecord returns true if the records matched, false otherwise
func checkRecord(w io.Writer, expected, actual map[string]any) bool {
	switch {
	case expected == nil && actual == nil:
		return true
	case expected == nil:
		actualJson, _ := json.Marshal(actual)
		_ = airbyte.LogError(w, fmt.Errorf("received unexpected record: %s", actualJson))
		return false
	case actual == nil:
		expectedJson, _ := json.Marshal(expected)
		_ = airbyte.LogError(w, fmt.Errorf("did not receive expected record: %s", expectedJson))
		return false
	}
	if d := diff.Diff(dataset.ConvertIntToFloat64(expected), actual); d != "" {
		_ = airbyte.LogError(w, fmt.Errorf("unexpected record (-expected, +received):\n%s", d))
		return false
	}
	return true
}
