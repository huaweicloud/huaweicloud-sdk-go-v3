package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateSpeedRequest struct {
	// 主机迁移任务的id

	TaskId string `json:"task_id"`

	Body *SpeedLimit `json:"body,omitempty"`
}

func (o UpdateSpeedRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateSpeedRequest struct{}"
	}

	return strings.Join([]string{"UpdateSpeedRequest", string(data)}, " ")
}
