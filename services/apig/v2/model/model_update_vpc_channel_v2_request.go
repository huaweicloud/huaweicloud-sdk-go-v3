package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateVpcChannelV2Request struct {
	// 实例编号

	InstanceId string `json:"instance_id"`
	// VPC通道的编号

	VpcChannelId string `json:"vpc_channel_id"`

	Body *VpcCreate `json:"body,omitempty"`
}

func (o UpdateVpcChannelV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateVpcChannelV2Request struct{}"
	}

	return strings.Join([]string{"UpdateVpcChannelV2Request", string(data)}, " ")
}
