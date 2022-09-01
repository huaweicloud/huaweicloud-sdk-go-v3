package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 正在召开研讨会议信息（含基本信息，会议UUID 及实时在线人数
type OpenWebinarOngoingInfo struct {

	// 实时在线人数
	OnlineAttendeeCount *int32 `json:"onlineAttendeeCount,omitempty" xml:"onlineAttendeeCount"`

	// 会议UUID
	ConfUUID *string `json:"confUUID,omitempty" xml:"confUUID"`

	// 预订人部门
	DeptName *string `json:"deptName,omitempty" xml:"deptName"`

	// 会议ID。长度限制为32个字符。
	ConferenceId *string `json:"conferenceId,omitempty" xml:"conferenceId"`

	// 企业id
	CorpId *string `json:"corpId,omitempty" xml:"corpId"`

	// 主题
	Subject *string `json:"subject,omitempty" xml:"subject"`

	// 描述
	Description *string `json:"description,omitempty" xml:"description"`

	// 会议召开时间
	StartTime *string `json:"startTime,omitempty" xml:"startTime"`

	// 时区ID
	TimeZoneId *int32 `json:"timeZoneId,omitempty" xml:"timeZoneId"`

	// 会议预订者ID
	ScheduserId *string `json:"scheduserId,omitempty" xml:"scheduserId"`

	// 会议预订者帐号名称。长度最大限制为96个字符。
	ScheduserName *string `json:"scheduserName,omitempty" xml:"scheduserName"`

	// 网络研讨会资源名
	VmrPkgName *string `json:"vmrPkgName,omitempty" xml:"vmrPkgName"`

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
}

func (o OpenWebinarOngoingInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpenWebinarOngoingInfo struct{}"
	}

	return strings.Join([]string{"OpenWebinarOngoingInfo", string(data)}, " ")
}
