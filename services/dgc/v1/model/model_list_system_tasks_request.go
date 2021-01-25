package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListSystemTasksRequest struct {
	TaskId string `json:"task_id"`
}

func (o ListSystemTasksRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListSystemTasksRequest struct{}"
	}

	return strings.Join([]string{"ListSystemTasksRequest", string(data)}, " ")
}
