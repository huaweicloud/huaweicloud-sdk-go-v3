package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteVpcChannelV2Request struct {
	// 实例ID

	InstanceId string `json:"instance_id"`
	// VPC通道的编号

	VpcChannelId string `json:"vpc_channel_id"`
}

func (o DeleteVpcChannelV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteVpcChannelV2Request struct{}"
	}

	return strings.Join([]string{"DeleteVpcChannelV2Request", string(data)}, " ")
}
