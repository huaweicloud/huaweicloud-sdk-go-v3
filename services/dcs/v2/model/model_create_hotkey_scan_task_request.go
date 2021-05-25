package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateHotkeyScanTaskRequest struct {
	// 实例ID。

	InstanceId string `json:"instance_id"`
}

func (o CreateHotkeyScanTaskRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateHotkeyScanTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateHotkeyScanTaskRequest", string(data)}, " ")
}
