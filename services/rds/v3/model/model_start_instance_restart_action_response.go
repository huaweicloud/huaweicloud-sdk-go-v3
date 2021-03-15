package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type StartInstanceRestartActionResponse struct {
	// 任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o StartInstanceRestartActionResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "StartInstanceRestartActionResponse struct{}"
	}

	return strings.Join([]string{"StartInstanceRestartActionResponse", string(data)}, " ")
}
