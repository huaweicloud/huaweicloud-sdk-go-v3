package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OpenWebinarBaseInfo 网络研讨会信息。
type OpenWebinarBaseInfo struct {

	// 网络研讨会ID。
	ConferenceId *string `json:"conferenceId,omitempty"`

	// 企业ID。
	CorpId *string `json:"corpId,omitempty"`

	// 网络研讨会主题。
	Subject *string `json:"subject,omitempty"`

	// 网络研讨会描述。
	Description *string `json:"description,omitempty"`

	// 网络研讨会开始时间（UTC时间），格式“yyyy-MM-dd HH:mm”。
	StartTime *string `json:"startTime,omitempty"`

	// 网络研讨会持续时长，单位分钟，取值范围[15,1440]。
	Duration *int32 `json:"duration,omitempty"`

	// 会议通知中会议时间的时区信息。时区信息，参考[[时区映射关系](https://support.huaweicloud.com/api-meeting/meeting_21_0110.html#ZH-CN_TOPIC_0212714472__table137407441463)](tag:hws)[[时区映射关系](https://support.huaweicloud.com/intl/zh-cn/api-meeting/meeting_21_0110.html#ZH-CN_TOPIC_0212714472__table137407441463)](tag:hk)。 > * 举例：“timeZoneID”:\"26\"，则通过华为云会议发送的会议通知中的时间将会标记为如“2021/11/11 星期四 00:00 - 02:00 (GMT) 格林威治标准时间:都柏林, 爱丁堡, 里斯本, 伦敦”。
	TimeZoneId *int32 `json:"timeZoneId,omitempty"`

	State *MeetingStatus `json:"state,omitempty"`

	// 网络研讨会预订者的用户UUID。
	ScheduserId *string `json:"scheduserId,omitempty"`

	// 预订者部门命名。
	DeptName *string `json:"deptName,omitempty"`

	// 预订者名称。
	ScheduserName *string `json:"scheduserName,omitempty"`

	// 网络研讨会VMR名称。
	VmrPkgName *string `json:"vmrPkgName,omitempty"`

	// 入会范围开关。
	CallRestriction *bool `json:"callRestriction,omitempty"`

	// 主持人、嘉宾入会范围。 * 0: 所有用户 * 2: 企业内用户 * 3: 被邀请用户
	Scope *int32 `json:"scope,omitempty"`

	// 观众入会范围。 * 0: 所有用户 * 2: 企业内用户
	AudienceScope *int32 `json:"audienceScope,omitempty"`

	// 主持人入会地址。
	ChairJoinUri *string `json:"chairJoinUri,omitempty"`

	// 主持人入会密码。
	ChairPasswd *string `json:"chairPasswd,omitempty"`

	// 嘉宾入会地址。
	GuestJoinUri *string `json:"guestJoinUri,omitempty"`

	// 嘉宾入会密码。
	GuestPasswd *string `json:"guestPasswd,omitempty"`

	// 观众入会地址。
	AudienceJoinUri *string `json:"audienceJoinUri,omitempty"`

	// 观众入会密码。
	AudiencePasswd *string `json:"audiencePasswd,omitempty"`

	EnableRecording *YesNoEnum `json:"enableRecording,omitempty"`
}

func (o OpenWebinarBaseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpenWebinarBaseInfo struct{}"
	}

	return strings.Join([]string{"OpenWebinarBaseInfo", string(data)}, " ")
}
