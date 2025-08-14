package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GroupMembershipItem 关联关系
type GroupMembershipItem struct {

	// 身份源中IdentityCenter用户组的全局唯一标识符（ID）
	GroupId *string `json:"group_id,omitempty"`

	// 身份源的全局唯一标识符（ID）
	IdentityStoreId *string `json:"identity_store_id,omitempty"`

	MemberId *MemberIdDto `json:"member_id,omitempty"`

	// 身份源中用户和组关联关系的全局唯一标识符（ID）
	MembershipId *string `json:"membership_id,omitempty"`
}

func (o GroupMembershipItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GroupMembershipItem struct{}"
	}

	return strings.Join([]string{"GroupMembershipItem", string(data)}, " ")
}
