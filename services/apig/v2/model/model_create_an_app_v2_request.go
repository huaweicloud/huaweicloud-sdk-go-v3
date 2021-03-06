package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateAnAppV2Request struct {
	// 实例编号

	InstanceId string `json:"instance_id"`

	Body *AppReq `json:"body,omitempty"`
}

func (o CreateAnAppV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateAnAppV2Request struct{}"
	}

	return strings.Join([]string{"CreateAnAppV2Request", string(data)}, " ")
}
