package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ExecuteQaChatRequest struct {
	// 机器人标识符，qabot编号，UUID格式。

	QabotId string `json:"qabot_id"`

	Body *PostRequestsReq `json:"body,omitempty"`
}

func (o ExecuteQaChatRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ExecuteQaChatRequest struct{}"
	}

	return strings.Join([]string{"ExecuteQaChatRequest", string(data)}, " ")
}
