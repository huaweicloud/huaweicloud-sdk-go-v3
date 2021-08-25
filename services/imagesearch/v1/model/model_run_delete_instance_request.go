package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type RunDeleteInstanceRequest struct {
	// 实例名称。

	InstanceName string `json:"instance_name"`
}

func (o RunDeleteInstanceRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RunDeleteInstanceRequest struct{}"
	}

	return strings.Join([]string{"RunDeleteInstanceRequest", string(data)}, " ")
}
