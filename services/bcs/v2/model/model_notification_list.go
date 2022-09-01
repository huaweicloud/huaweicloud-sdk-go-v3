package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NotificationList struct {

	// 通道名称
	ChannelName *string `json:"channel_name,omitempty" xml:"channel_name"`

	// 当前状态
	Status *string `json:"status,omitempty" xml:"status"`

	// 更新时间
	UpdatedTime *string `json:"updated_time,omitempty" xml:"updated_time"`

	// 是否开启可信
	Tc3Need *bool `json:"tc3_need,omitempty" xml:"tc3_need"`

	InvitorInfo *InvitorInfo `json:"invitor_info,omitempty" xml:"invitor_info"`

	InviteeInfo *InviteeInfo `json:"invitee_info,omitempty" xml:"invitee_info"`

	// 是否删除
	Hide *int64 `json:"hide,omitempty" xml:"hide"`

	// 被邀请的组织信息
	InviteeOrgs *[]OrganizationV2 `json:"invitee_orgs,omitempty" xml:"invitee_orgs"`

	// 阅读状态值
	ReadStatus *int64 `json:"read_status,omitempty" xml:"read_status"`

	// 跨版本进行升级
	CrossVersionUpgrade *string `json:"cross_version_upgrade,omitempty" xml:"cross_version_upgrade"`
}

func (o NotificationList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NotificationList struct{}"
	}

	return strings.Join([]string{"NotificationList", string(data)}, " ")
}
