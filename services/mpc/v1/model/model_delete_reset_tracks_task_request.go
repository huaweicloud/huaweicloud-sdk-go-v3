package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteResetTracksTaskRequest struct {
	// 任务ID

	TaskId string `json:"task_id"`
}

func (o DeleteResetTracksTaskRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteResetTracksTaskRequest struct{}"
	}

	return strings.Join([]string{"DeleteResetTracksTaskRequest", string(data)}, " ")
}
