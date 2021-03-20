package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateEnvironmentV2Request struct {
	InstanceId string `json:"instance_id"`

	Body *EnvReq `json:"body,omitempty"`
}

func (o CreateEnvironmentV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateEnvironmentV2Request struct{}"
	}

	return strings.Join([]string{"CreateEnvironmentV2Request", string(data)}, " ")
}
