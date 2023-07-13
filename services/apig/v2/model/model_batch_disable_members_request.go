package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDisableMembersRequest Request Object
type BatchDisableMembersRequest struct {

	// 实例ID，在API网关控制台的“实例信息”中获取。
	InstanceId string `json:"instance_id"`

	// VPC通道的编号
	VpcChannelId string `json:"vpc_channel_id"`

	Body *MembersBatchEnableOrDisable `json:"body,omitempty"`
}

func (o BatchDisableMembersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDisableMembersRequest struct{}"
	}

	return strings.Join([]string{"BatchDisableMembersRequest", string(data)}, " ")
}
