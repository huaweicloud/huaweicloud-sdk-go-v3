package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 联盟成员
type Member struct {
	// 是否支持可信

	TcsNeed *bool `json:"tcs_need,omitempty"`
	// 通道名称

	ChannelName *string `json:"channel_name,omitempty"`
	// 被邀请的组织

	InvitedOrgs *[]OrganizationV2 `json:"invited_orgs,omitempty"`

	InvitorInfo *MemberInvitor `json:"invitor_info,omitempty"`

	InviteeInfo *MemberInvitee `json:"invitee_info,omitempty"`
}

func (o Member) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Member struct{}"
	}

	return strings.Join([]string{"Member", string(data)}, " ")
}
