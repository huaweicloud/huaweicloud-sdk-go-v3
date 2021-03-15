package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type StartInstanceSlaveDrActionResponse struct {
	// 任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o StartInstanceSlaveDrActionResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "StartInstanceSlaveDrActionResponse struct{}"
	}

	return strings.Join([]string{"StartInstanceSlaveDrActionResponse", string(data)}, " ")
}
