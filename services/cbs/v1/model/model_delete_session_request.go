package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteSessionRequest struct {
	// 机器人标识符。

	QabotId string `json:"qabot_id"`
	// 会话标识符。

	SessionId string `json:"session_id"`
}

func (o DeleteSessionRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteSessionRequest struct{}"
	}

	return strings.Join([]string{"DeleteSessionRequest", string(data)}, " ")
}
