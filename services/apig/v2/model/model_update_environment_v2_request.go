package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateEnvironmentV2Request struct {
	InstanceId string `json:"instance_id"`

	EnvId string `json:"env_id"`

	Body *EnvReq `json:"body,omitempty"`
}

func (o UpdateEnvironmentV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateEnvironmentV2Request struct{}"
	}

	return strings.Join([]string{"UpdateEnvironmentV2Request", string(data)}, " ")
}
