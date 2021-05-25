package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteBigkeyScanTaskRequest struct {
	// 实例ID。

	InstanceId string `json:"instance_id"`
	// 大key分析任务ID。

	BigkeyId string `json:"bigkey_id"`
}

func (o DeleteBigkeyScanTaskRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteBigkeyScanTaskRequest struct{}"
	}

	return strings.Join([]string{"DeleteBigkeyScanTaskRequest", string(data)}, " ")
}
