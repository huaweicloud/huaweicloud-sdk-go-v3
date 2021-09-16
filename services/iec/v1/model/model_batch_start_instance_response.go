package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type BatchStartInstanceResponse struct {
	// 任务列表对象。

	Jobs           *[]JobResult `json:"jobs,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o BatchStartInstanceResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchStartInstanceResponse struct{}"
	}

	return strings.Join([]string{"BatchStartInstanceResponse", string(data)}, " ")
}
