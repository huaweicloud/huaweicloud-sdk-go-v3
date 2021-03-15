package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListBatchTasksRequest struct {
	InstanceId *string `json:"Instance-Id,omitempty"`
	AppId      *string `json:"app_id,omitempty"`
	TaskType   string  `json:"task_type"`
	Status     *string `json:"status,omitempty"`
	Limit      *int32  `json:"limit,omitempty"`
	Marker     *string `json:"marker,omitempty"`
	Offset     *int32  `json:"offset,omitempty"`
}

func (o ListBatchTasksRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListBatchTasksRequest struct{}"
	}

	return strings.Join([]string{"ListBatchTasksRequest", string(data)}, " ")
}
