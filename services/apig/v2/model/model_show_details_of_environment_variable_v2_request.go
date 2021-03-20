package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowDetailsOfEnvironmentVariableV2Request struct {
	InstanceId string `json:"instance_id"`

	EnvVariableId string `json:"env_variable_id"`
}

func (o ShowDetailsOfEnvironmentVariableV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowDetailsOfEnvironmentVariableV2Request struct{}"
	}

	return strings.Join([]string{"ShowDetailsOfEnvironmentVariableV2Request", string(data)}, " ")
}
