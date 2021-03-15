package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowBatchTaskRequest struct {
	InstanceId *string `json:"Instance-Id,omitempty"`
	TaskId     string  `json:"task_id"`
	Limit      *int32  `json:"limit,omitempty"`
	Marker     *string `json:"marker,omitempty"`
	Offset     *int32  `json:"offset,omitempty"`
}

func (o ShowBatchTaskRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowBatchTaskRequest struct{}"
	}

	return strings.Join([]string{"ShowBatchTaskRequest", string(data)}, " ")
}
