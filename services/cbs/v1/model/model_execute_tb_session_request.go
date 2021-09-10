package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ExecuteTbSessionRequest struct {
	// 话务机器人ID

	BotId string `json:"bot_id"`
	// 会话ID，在发起会话API的请求结果消息体response当中产生。

	SessionId string `json:"session_id"`

	Body *ExecuteTbSessionReq `json:"body,omitempty"`
}

func (o ExecuteTbSessionRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ExecuteTbSessionRequest struct{}"
	}

	return strings.Join([]string{"ExecuteTbSessionRequest", string(data)}, " ")
}
