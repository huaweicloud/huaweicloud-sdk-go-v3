package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 联盟成员
type Member struct {

	// 是否支持可信
	TcsNeed *bool `json:"tcs_need,omitempty" xml:"tcs_need"`

	// 通道名称
	ChannelName *string `json:"channel_name,omitempty" xml:"channel_name"`

	// 被邀请的组织
	InvitedOrgs *[]OrganizationV2 `json:"invited_orgs,omitempty" xml:"invited_orgs"`

	InvitorInfo *MemberInvitor `json:"invitor_info,omitempty" xml:"invitor_info"`

	InviteeInfo *MemberInvitee `json:"invitee_info,omitempty" xml:"invitee_info"`
}

func (o Member) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Member struct{}"
	}

	return strings.Join([]string{"Member", string(data)}, " ")
}
