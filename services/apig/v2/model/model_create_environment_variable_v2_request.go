package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateEnvironmentVariableV2Request struct {
	InstanceId string          `json:"instance_id"`
	Body       *EnvVariableReq `json:"body,omitempty"`
}

func (o CreateEnvironmentVariableV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateEnvironmentVariableV2Request struct{}"
	}

	return strings.Join([]string{"CreateEnvironmentVariableV2Request", string(data)}, " ")
}
