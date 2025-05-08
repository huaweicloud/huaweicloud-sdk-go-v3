package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GroupMemberEntity 成员详情
type GroupMemberEntity struct {

	// 成员Id
	MemberId *string `json:"member_id,omitempty"`

	// 客户端Id
	ClientId *string `json:"client_id,omitempty"`
}

func (o GroupMemberEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GroupMemberEntity struct{}"
	}

	return strings.Join([]string{"GroupMemberEntity", string(data)}, " ")
}
