package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteEnvironmentVariableV2Request struct {
	InstanceId    string `json:"instance_id"`
	EnvVariableId string `json:"env_variable_id"`
}

func (o DeleteEnvironmentVariableV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteEnvironmentVariableV2Request struct{}"
	}

	return strings.Join([]string{"DeleteEnvironmentVariableV2Request", string(data)}, " ")
}
