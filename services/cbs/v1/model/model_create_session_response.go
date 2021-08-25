package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateSessionResponse struct {
	// 会话标识符。

	SessionId *string `json:"session_id,omitempty"`
	// 机器人问候语。

	Greeting       *string `json:"greeting,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateSessionResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateSessionResponse struct{}"
	}

	return strings.Join([]string{"CreateSessionResponse", string(data)}, " ")
}
