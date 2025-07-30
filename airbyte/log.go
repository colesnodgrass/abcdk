package airbyte

import (
	"encoding/json"
	"fmt"
	"github.com/colesnodgrass/abcdk/protocol"
	"io"
)

func toMsg(level protocol.AirbyteLogMessageLevel, msg string) protocol.AirbyteMessage {
	msgLog := protocol.AirbyteLogMessage{
		Level:   level,
		Message: msg,
	}

	return protocol.AirbyteMessage{
		Log:  &msgLog,
		Type: protocol.AirbyteMessageTypeLOG,
	}
}

func LogInfo(w io.Writer, msg string) error {
	msgLog := toMsg(protocol.AirbyteLogMessageLevelINFO, msg)
	data, err := json.Marshal(msgLog)
	if err != nil {
		return fmt.Errorf("failed to marshal log: %w", err)
	}
	_, err = w.Write(data)
	return err
}

func LogError(w io.Writer, err error) error {
	msgLog := toMsg(protocol.AirbyteLogMessageLevelERROR, err.Error())
	data, err := json.Marshal(msgLog)
	if err != nil {
		return fmt.Errorf("failed to marshal log: %w", err)
	}
	_, err = w.Write(data)
	return err
}
