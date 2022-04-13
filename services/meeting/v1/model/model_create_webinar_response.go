package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateWebinarResponse struct {
	// 会议ID。长度限制为32个字符。

	ConferenceId *string `json:"conferenceId,omitempty"`
	// 企业id

	CorpId *string `json:"corpId,omitempty"`
	// 主题

	Subject *string `json:"subject,omitempty"`
	// 描述

	Description *string `json:"description,omitempty"`
	// 开始时间

	StartTime *string `json:"startTime,omitempty"`
	// 时长，单位分钟

	Duration *int32 `json:"duration,omitempty"`
	// 时区ID

	TimeZoneId *int32 `json:"timeZoneId,omitempty"`

	State *MeetingStatus `json:"state,omitempty"`
	// 会议预订者ID

	ScheduserId *string `json:"scheduserId,omitempty"`
	// 预订人部门

	DeptName *string `json:"deptName,omitempty"`
	// 会议预订者帐号名称。长度最大限制为96个字符。

	ScheduserName *string `json:"scheduserName,omitempty"`
	// 网络研讨会资源名

	VmrPkgName *string `json:"vmrPkgName,omitempty"`
	// 入会范围开关

	CallRestriction *bool `json:"callRestriction,omitempty"`
	// 主持人、嘉宾入会范围  0: 所有用户 1: 非匿名用户（手机pstn入会视为匿名入会） 2: 企业内用户 3: 被邀请用户。

	Scope *int32 `json:"scope,omitempty"`
	// 观众入会范围 0: 所有用户 2: 企业内用户和被邀请用户。

	AudienceScope *int32 `json:"audienceScope,omitempty"`
	// 主持人入会地址。

	ChairJoinUri *string `json:"chairJoinUri,omitempty"`
	// 主持人密码。

	ChairPasswd *string `json:"chairPasswd,omitempty"`
	// 嘉宾入会地址。

	GuestJoinUri *string `json:"guestJoinUri,omitempty"`
	// 嘉宾密码。

	GuestPasswd *string `json:"guestPasswd,omitempty"`
	// 观众入会地址。

	AudienceJoinUri *string `json:"audienceJoinUri,omitempty"`
	// 观众密码。

	AudiencePasswd *string `json:"audiencePasswd,omitempty"`

	Attendees      *[]string `json:"attendees,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o CreateWebinarResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateWebinarResponse struct{}"
	}

	return strings.Join([]string{"CreateWebinarResponse", string(data)}, " ")
}
