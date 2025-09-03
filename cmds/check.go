package cmds

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/colesnodgrass/abcdk/config"
	"github.com/colesnodgrass/abcdk/protocol"
)

type CheckCmd struct {
	Config string `type:"existingfile"`
}

func (cc *CheckCmd) Run(ctx context.Context, w io.Writer) error {
	cfg, err := config.FromFile(cc.Config)
	if err != nil {
		return fmt.Errorf("failed to load config %s: %w", cc.Config, err)
	}

	var response protocol.AirbyteMessage
	switch cfg.Check {
	case "pass":
		response = cc.msg(cc.pass())
	case "fail":
		response = cc.msg(cc.fail())
	default:
		return fmt.Errorf("unsupported config dat for check %s", cfg.Check)
	}

	data, err := json.Marshal(response)
	if err != nil {
		return fmt.Errorf("failed to marshal status: %w", err)
	}
	if _, err := w.Write(data); err != nil {
		return fmt.Errorf("failed to write: %w\n", err)
	}

	return nil
}

func (cc *CheckCmd) msg(status protocol.AirbyteConnectionStatus) protocol.AirbyteMessage {
	return protocol.AirbyteMessage{
		ConnectionStatus: &status,
		Type:             protocol.AirbyteMessageTypeCONNECTIONSTATUS,
	}
}

func (cc *CheckCmd) pass() protocol.AirbyteConnectionStatus {
	return protocol.AirbyteConnectionStatus{
		Message: p("check passed (expected)"),
		Status:  protocol.AirbyteConnectionStatusStatusSUCCEEDED,
	}
}

func (cc *CheckCmd) fail() protocol.AirbyteConnectionStatus {
	return protocol.AirbyteConnectionStatus{
		Message: p("check failed (expected)"),
		Status:  protocol.AirbyteConnectionStatusStatusFAILED,
	}
}
