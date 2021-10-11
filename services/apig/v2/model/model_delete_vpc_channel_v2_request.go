package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteVpcChannelV2Request struct {
	// 实例编号

	InstanceId string `json:"instance_id"`
	// VPC通道的编号

	VpcChannelId string `json:"vpc_channel_id"`
}

func (o DeleteVpcChannelV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteVpcChannelV2Request struct{}"
	}

	return strings.Join([]string{"DeleteVpcChannelV2Request", string(data)}, " ")
}
