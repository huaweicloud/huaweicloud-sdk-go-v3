package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateVpcChannelV2Request struct {
	// 实例编号

	InstanceId string `json:"instance_id"`

	Body *VpcCreate `json:"body,omitempty"`
}

func (o CreateVpcChannelV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateVpcChannelV2Request struct{}"
	}

	return strings.Join([]string{"CreateVpcChannelV2Request", string(data)}, " ")
}
