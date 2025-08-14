package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GroupMembershipExistenceResultDto Indicates whether a resource is a member of a group in the identity store.
type GroupMembershipExistenceResultDto struct {

	// 身份源中IdentityCenter用户组的全局唯一标识符（ID）
	GroupId *string `json:"group_id,omitempty"`

	MemberId *MemberIdDto `json:"member_id,omitempty"`

	// 一个布尔值，表示用户是否在组中
	MembershipExists *bool `json:"membership_exists,omitempty"`
}

func (o GroupMembershipExistenceResultDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GroupMembershipExistenceResultDto struct{}"
	}

	return strings.Join([]string{"GroupMembershipExistenceResultDto", string(data)}, " ")
}
