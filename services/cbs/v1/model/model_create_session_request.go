package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateSessionRequest struct {
	// 机器人标识符。

	QabotId string `json:"qabot_id"`
}

func (o CreateSessionRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateSessionRequest struct{}"
	}

	return strings.Join([]string{"CreateSessionRequest", string(data)}, " ")
}
