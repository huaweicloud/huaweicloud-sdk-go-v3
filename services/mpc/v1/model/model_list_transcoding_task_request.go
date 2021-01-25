package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListTranscodingTaskRequest struct {
	XLanguage *string  `json:"x-language,omitempty"`
	TaskId    *[]int64 `json:"task_id,omitempty"`
	Status    *string  `json:"status,omitempty"`
	StartTime *string  `json:"start_time,omitempty"`
	EndTime   *string  `json:"end_time,omitempty"`
	Page      *int32   `json:"page,omitempty"`
	Size      *int32   `json:"size,omitempty"`
}

func (o ListTranscodingTaskRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListTranscodingTaskRequest struct{}"
	}

	return strings.Join([]string{"ListTranscodingTaskRequest", string(data)}, " ")
}
