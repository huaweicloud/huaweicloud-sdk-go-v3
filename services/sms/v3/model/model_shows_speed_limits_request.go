package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowsSpeedLimitsRequest struct {
	// 查询限速信息的任务id

	TaskId string `json:"task_id"`
}

func (o ShowsSpeedLimitsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowsSpeedLimitsRequest struct{}"
	}

	return strings.Join([]string{"ShowsSpeedLimitsRequest", string(data)}, " ")
}
