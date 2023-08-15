package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateBackendInstancesV2Request Request Object
type UpdateBackendInstancesV2Request struct {

	// 实例ID，在API网关控制台的“实例信息”中获取。
	InstanceId string `json:"instance_id"`

	// VPC通道的编号
	VpcChannelId string `json:"vpc_channel_id"`

	Body *VpcMemberModify `json:"body,omitempty"`
}

func (o UpdateBackendInstancesV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateBackendInstancesV2Request struct{}"
	}

	return strings.Join([]string{"UpdateBackendInstancesV2Request", string(data)}, " ")
}
