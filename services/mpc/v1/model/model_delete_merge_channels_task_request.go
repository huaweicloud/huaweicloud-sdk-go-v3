package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteMergeChannelsTaskRequest struct {
	// 任务ID

	TaskId string `json:"task_id"`
}

func (o DeleteMergeChannelsTaskRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteMergeChannelsTaskRequest struct{}"
	}

	return strings.Join([]string{"DeleteMergeChannelsTaskRequest", string(data)}, " ")
}
