package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteEnvironmentV2Request struct {
	// 实例编号

	InstanceId string `json:"instance_id"`
	// 环境的ID

	EnvId string `json:"env_id"`
}

func (o DeleteEnvironmentV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteEnvironmentV2Request struct{}"
	}

	return strings.Join([]string{"DeleteEnvironmentV2Request", string(data)}, " ")
}
