package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateApiV2Request struct {
	// 实例编号

	InstanceId string `json:"instance_id"`

	Body *ApiCreate `json:"body,omitempty"`
}

func (o CreateApiV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateApiV2Request struct{}"
	}

	return strings.Join([]string{"CreateApiV2Request", string(data)}, " ")
}
