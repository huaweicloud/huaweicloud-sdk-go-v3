package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteHotkeyScanTaskRequest struct {
	// 实例ID。

	InstanceId string `json:"instance_id"`
	// 热key分析任务ID。

	HotkeyId string `json:"hotkey_id"`
}

func (o DeleteHotkeyScanTaskRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteHotkeyScanTaskRequest struct{}"
	}

	return strings.Join([]string{"DeleteHotkeyScanTaskRequest", string(data)}, " ")
}
