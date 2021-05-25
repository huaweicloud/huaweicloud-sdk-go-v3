package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteEnvironmentVariableV2Request struct {
	// 实例编号

	InstanceId string `json:"instance_id"`
	// 环境变量的ID

	EnvVariableId string `json:"env_variable_id"`
}

func (o DeleteEnvironmentVariableV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteEnvironmentVariableV2Request struct{}"
	}

	return strings.Join([]string{"DeleteEnvironmentVariableV2Request", string(data)}, " ")
}
