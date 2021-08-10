package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListTakeOverTaskResponse struct {
	// 托管任务信息。

	Tasks *[]TakeOverTask `json:"tasks,omitempty"`
	// 任务数量。

	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListTakeOverTaskResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListTakeOverTaskResponse struct{}"
	}

	return strings.Join([]string{"ListTakeOverTaskResponse", string(data)}, " ")
}
