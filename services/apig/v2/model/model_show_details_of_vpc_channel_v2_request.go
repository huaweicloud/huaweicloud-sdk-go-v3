package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowDetailsOfVpcChannelV2Request struct {
	// 实例编号

	InstanceId string `json:"instance_id"`
	// VPC通道的编号

	VpcChannelId string `json:"vpc_channel_id"`
}

func (o ShowDetailsOfVpcChannelV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowDetailsOfVpcChannelV2Request struct{}"
	}

	return strings.Join([]string{"ShowDetailsOfVpcChannelV2Request", string(data)}, " ")
}
