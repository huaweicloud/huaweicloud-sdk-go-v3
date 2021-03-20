package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type StartInstanceSingleToHaActionResponse struct {
	// 任务ID。

	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o StartInstanceSingleToHaActionResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "StartInstanceSingleToHaActionResponse struct{}"
	}

	return strings.Join([]string{"StartInstanceSingleToHaActionResponse", string(data)}, " ")
}
