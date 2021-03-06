package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type StopInstanceRequest struct {
	// 实例id

	InstanceId string `json:"instance_id"`
}

func (o StopInstanceRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "StopInstanceRequest struct{}"
	}

	return strings.Join([]string{"StopInstanceRequest", string(data)}, " ")
}
