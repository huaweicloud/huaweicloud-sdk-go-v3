package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateVpcChannelV2Request Request Object
type UpdateVpcChannelV2Request struct {

	// 实例ID，在API网关控制台的“实例信息”中获取。
	InstanceId string `json:"instance_id"`

	// VPC通道的编号
	VpcChannelId string `json:"vpc_channel_id"`

	Body *VpcCreate `json:"body,omitempty"`
}

func (o UpdateVpcChannelV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVpcChannelV2Request struct{}"
	}

	return strings.Join([]string{"UpdateVpcChannelV2Request", string(data)}, " ")
}
