package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateMemberGroupRequest struct {
	// 实例ID

	InstanceId string `json:"instance_id"`
	// VPC通道的编号

	VpcChannelId string `json:"vpc_channel_id"`

	Body *MemberGroupCreateBatch `json:"body,omitempty"`
}

func (o CreateMemberGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMemberGroupRequest struct{}"
	}

	return strings.Join([]string{"CreateMemberGroupRequest", string(data)}, " ")
}
