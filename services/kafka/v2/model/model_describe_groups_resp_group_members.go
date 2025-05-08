package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DescribeGroupsRespGroupMembers struct {

	// 消费组consumer地址。
	Host *string `json:"host,omitempty"`

	// 消费组consumer的ID。
	MemberId *string `json:"member_id,omitempty"`

	// 客户端ID。
	ClientId *string `json:"client_id,omitempty"`
}

func (o DescribeGroupsRespGroupMembers) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DescribeGroupsRespGroupMembers struct{}"
	}

	return strings.Join([]string{"DescribeGroupsRespGroupMembers", string(data)}, " ")
}
