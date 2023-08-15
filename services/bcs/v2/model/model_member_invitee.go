package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MemberInvitee 联盟成员中的被邀请方
type MemberInvitee struct {

	// 被邀请方BCS服务实例ID
	InviteeBcsId *string `json:"invitee_bcs_id,omitempty"`

	// 被邀请方租户id
	InviteeUserId *string `json:"invitee_user_id,omitempty"`

	// 被邀请方租户名
	InviteeUsername *string `json:"invitee_username,omitempty"`
}

func (o MemberInvitee) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MemberInvitee struct{}"
	}

	return strings.Join([]string{"MemberInvitee", string(data)}, " ")
}
