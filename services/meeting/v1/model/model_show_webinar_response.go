package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowWebinarResponse struct {

	// 会议ID。长度限制为32个字符。
	ConferenceId *string `json:"conferenceId,omitempty" xml:"conferenceId"`

	// 企业id
	CorpId *string `json:"corpId,omitempty" xml:"corpId"`

	// 主题
	Subject *string `json:"subject,omitempty" xml:"subject"`

	// 描述
	Description *string `json:"description,omitempty" xml:"description"`

	// 开始时间
	StartTime *string `json:"startTime,omitempty" xml:"startTime"`

	// 时长，单位分钟
	Duration *int32 `json:"duration,omitempty" xml:"duration"`

	// 时区ID
	TimeZoneId *int32 `json:"timeZoneId,omitempty" xml:"timeZoneId"`

	State *MeetingStatus `json:"state,omitempty" xml:"state"`

	// 会议预订者ID
	ScheduserId *string `json:"scheduserId,omitempty" xml:"scheduserId"`

	// 预订人部门
	DeptName *string `json:"deptName,omitempty" xml:"deptName"`

	// 会议预订者帐号名称。长度最大限制为96个字符。
	ScheduserName *string `json:"scheduserName,omitempty" xml:"scheduserName"`

	// 网络研讨会资源名
	VmrPkgName *string `json:"vmrPkgName,omitempty" xml:"vmrPkgName"`

	// 入会范围开关
	CallRestriction *bool `json:"callRestriction,omitempty" xml:"callRestriction"`

	// 主持人、嘉宾入会范围  0: 所有用户 1: 非匿名用户（手机pstn入会视为匿名入会） 2: 企业内用户 3: 被邀请用户。
	Scope *int32 `json:"scope,omitempty" xml:"scope"`

	// 观众入会范围 0: 所有用户 2: 企业内用户和被邀请用户。
	AudienceScope *int32 `json:"audienceScope,omitempty" xml:"audienceScope"`

	// 主持人入会地址。
	ChairJoinUri *string `json:"chairJoinUri,omitempty" xml:"chairJoinUri"`

	// 主持人密码。
	ChairPasswd *string `json:"chairPasswd,omitempty" xml:"chairPasswd"`

	// 嘉宾入会地址。
	GuestJoinUri *string `json:"guestJoinUri,omitempty" xml:"guestJoinUri"`

	// 嘉宾密码。
	GuestPasswd *string `json:"guestPasswd,omitempty" xml:"guestPasswd"`

	// 观众入会地址。
	AudienceJoinUri *string `json:"audienceJoinUri,omitempty" xml:"audienceJoinUri"`

	// 观众密码。
	AudiencePasswd *string `json:"audiencePasswd,omitempty" xml:"audiencePasswd"`

	NotifySetting *OpenNotifySetting `json:"notifySetting,omitempty" xml:"notifySetting"`

	Attendees      *[]string `json:"attendees,omitempty" xml:"attendees"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowWebinarResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWebinarResponse struct{}"
	}

	return strings.Join([]string{"ShowWebinarResponse", string(data)}, " ")
}
