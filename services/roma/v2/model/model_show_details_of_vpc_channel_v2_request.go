package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowDetailsOfVpcChannelV2Request struct {

	// 实例ID
	InstanceId string `json:"instance_id" xml:"instance_id"`

	// VPC通道的编号
	VpcChannelId string `json:"vpc_channel_id" xml:"vpc_channel_id"`
}

func (o ShowDetailsOfVpcChannelV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDetailsOfVpcChannelV2Request struct{}"
	}

	return strings.Join([]string{"ShowDetailsOfVpcChannelV2Request", string(data)}, " ")
}
