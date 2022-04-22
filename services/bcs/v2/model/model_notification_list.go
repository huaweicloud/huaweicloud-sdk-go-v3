package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NotificationList struct {

	// 通道名称
	ChannelName *string `json:"channel_name,omitempty"`

	// 当前状态
	Status *string `json:"status,omitempty"`

	// 更新时间
	UpdatedTime *string `json:"updated_time,omitempty"`

	// 是否开启可信
	Tc3Need *bool `json:"tc3_need,omitempty"`

	InvitorInfo *InvitorInfo `json:"invitor_info,omitempty"`

	InviteeInfo *InviteeInfo `json:"invitee_info,omitempty"`

	// 是否删除
	Hide *int64 `json:"hide,omitempty"`

	// 被邀请的组织信息
	InviteeOrgs *[]OrganizationV2 `json:"invitee_orgs,omitempty"`

	// 阅读状态值
	ReadStatus *int64 `json:"read_status,omitempty"`

	// 跨版本进行升级
	CrossVersionUpgrade *string `json:"cross_version_upgrade,omitempty"`
}

func (o NotificationList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NotificationList struct{}"
	}

	return strings.Join([]string{"NotificationList", string(data)}, " ")
}
