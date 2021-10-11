package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateEngressEipV2Request struct {
	// 实例编号

	InstanceId string `json:"instance_id"`

	Body *OpenEngressEipReq `json:"body,omitempty"`
}

func (o UpdateEngressEipV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateEngressEipV2Request struct{}"
	}

	return strings.Join([]string{"UpdateEngressEipV2Request", string(data)}, " ")
}
