package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type BatchStopInstanceResponse struct {
	// 任务列表对象。

	Jobs           *[]JobResult `json:"jobs,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o BatchStopInstanceResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchStopInstanceResponse struct{}"
	}

	return strings.Join([]string{"BatchStopInstanceResponse", string(data)}, " ")
}
