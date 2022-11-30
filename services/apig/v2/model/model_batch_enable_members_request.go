package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type BatchEnableMembersRequest struct {

	// 实例ID，在API网关控制台的“实例信息”中获取。
	InstanceId string `json:"instance_id"`

	// VPC通道的编号
	VpcChannelId string `json:"vpc_channel_id"`

	Body *MembersBatchEnableOrDisable `json:"body,omitempty"`
}

func (o BatchEnableMembersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchEnableMembersRequest struct{}"
	}

	return strings.Join([]string{"BatchEnableMembersRequest", string(data)}, " ")
}
