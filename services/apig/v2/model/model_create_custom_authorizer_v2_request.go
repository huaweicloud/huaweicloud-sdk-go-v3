package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateCustomAuthorizerV2Request struct {
	// 实例编号

	InstanceId string `json:"instance_id"`

	Body *AuthorizerCreate `json:"body,omitempty"`
}

func (o CreateCustomAuthorizerV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateCustomAuthorizerV2Request struct{}"
	}

	return strings.Join([]string{"CreateCustomAuthorizerV2Request", string(data)}, " ")
}
