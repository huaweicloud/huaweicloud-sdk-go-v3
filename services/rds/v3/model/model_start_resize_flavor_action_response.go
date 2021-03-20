package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type StartResizeFlavorActionResponse struct {
	// 任务ID。

	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o StartResizeFlavorActionResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "StartResizeFlavorActionResponse struct{}"
	}

	return strings.Join([]string{"StartResizeFlavorActionResponse", string(data)}, " ")
}
