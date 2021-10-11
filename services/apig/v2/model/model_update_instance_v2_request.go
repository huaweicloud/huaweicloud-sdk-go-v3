package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateInstanceV2Request struct {
	// 实例编号

	InstanceId string `json:"instance_id"`

	Body *InstanceModReq `json:"body,omitempty"`
}

func (o UpdateInstanceV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateInstanceV2Request struct{}"
	}

	return strings.Join([]string{"UpdateInstanceV2Request", string(data)}, " ")
}
