package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type BatchDisableMembersRequest struct {
	// 实例ID

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
