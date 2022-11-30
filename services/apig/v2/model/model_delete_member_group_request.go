package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteMemberGroupRequest struct {

	// 实例ID，在API网关控制台的“实例信息”中获取。
	InstanceId string `json:"instance_id"`

	// VPC通道的编号
	VpcChannelId string `json:"vpc_channel_id"`

	// VPC通道后端服务器组编号
	MemberGroupId string `json:"member_group_id"`
}

func (o DeleteMemberGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteMemberGroupRequest struct{}"
	}

	return strings.Join([]string{"DeleteMemberGroupRequest", string(data)}, " ")
}
