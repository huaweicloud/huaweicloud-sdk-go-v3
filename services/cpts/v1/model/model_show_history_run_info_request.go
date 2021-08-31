package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowHistoryRunInfoRequest struct {
	// 任务id

	TaskId int32 `json:"task_id"`
}

func (o ShowHistoryRunInfoRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowHistoryRunInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowHistoryRunInfoRequest", string(data)}, " ")
}
