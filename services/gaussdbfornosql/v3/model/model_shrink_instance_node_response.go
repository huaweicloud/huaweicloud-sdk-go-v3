package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShrinkInstanceNodeResponse struct {
	// 任务ID。

	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShrinkInstanceNodeResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShrinkInstanceNodeResponse struct{}"
	}

	return strings.Join([]string{"ShrinkInstanceNodeResponse", string(data)}, " ")
}
