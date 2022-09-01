package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowGroupsRespGroupMembers struct {

	// 消费组consumer地址。
	Host *string `json:"host,omitempty" xml:"host"`

	// consumer分配到的分区信息。
	Assignment *[]ShowGroupsRespGroupAssignment `json:"assignment,omitempty" xml:"assignment"`

	// 消费组consumer的ID。
	MemberId *string `json:"member_id,omitempty" xml:"member_id"`

	// 客户端ID。
	ClientId *string `json:"client_id,omitempty" xml:"client_id"`
}

func (o ShowGroupsRespGroupMembers) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGroupsRespGroupMembers struct{}"
	}

	return strings.Join([]string{"ShowGroupsRespGroupMembers", string(data)}, " ")
}
