package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type RunQueryInstanceRequest struct {
	// 实例名称。

	InstanceName string `json:"instance_name"`
}

func (o RunQueryInstanceRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RunQueryInstanceRequest struct{}"
	}

	return strings.Join([]string{"RunQueryInstanceRequest", string(data)}, " ")
}
