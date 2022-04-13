package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowDetailsOfMemberGroupRequest struct {
	// 实例ID

	InstanceId string `json:"instance_id"`
	// VPC通道的编号

	VpcChannelId string `json:"vpc_channel_id"`
	// VPC通道后端服务器组编号

	MemberGroupId string `json:"member_group_id"`
}

func (o ShowDetailsOfMemberGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDetailsOfMemberGroupRequest struct{}"
	}

	return strings.Join([]string{"ShowDetailsOfMemberGroupRequest", string(data)}, " ")
}
