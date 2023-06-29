package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OpenWebinarHistoryInfo 历史召开研讨会议信息。
type OpenWebinarHistoryInfo struct {

	// 网络研讨会ID。
	ConferenceId *string `json:"conferenceId,omitempty"`

	// 网络研讨会UUID。
	ConfUUID *string `json:"confUUID,omitempty"`

	// 网络研讨会主题。
	Subject *string `json:"subject,omitempty"`

	// 网络研讨会预定者名称。
	ScheduserName *string `json:"scheduserName,omitempty"`

	// 网络研讨主持人名称。
	Moderator *string `json:"moderator,omitempty"`

	// 预订人部门名称。
	DeptName *string `json:"deptName,omitempty"`

	// 会议通知中会议时间的时区信息。时区信息，参考[[时区映射关系](https://support.huaweicloud.com/api-meeting/meeting_21_0110.html#ZH-CN_TOPIC_0212714472__table137407441463)](tag:hws)[[时区映射关系](https://support.huaweicloud.com/intl/zh-cn/api-meeting/meeting_21_0110.html#ZH-CN_TOPIC_0212714472__table137407441463)](tag:hk)。 > * 举例：“timeZoneID”:\"26\"，则通过华为云会议发送的会议通知中的时间将会标记为如“2021/11/11 星期四 00:00 - 02:00 (GMT) 格林威治标准时间:都柏林, 爱丁堡, 里斯本, 伦敦”。
	TimeZoneId *int32 `json:"timeZoneId,omitempty"`

	// 网络研讨会开始时间（UTC时间），格式“yyyy-MM-dd HH:mm”。
	StartTime *string `json:"startTime,omitempty"`

	// 网络研讨会持续时长，单位分钟。
	Duration *int32 `json:"duration,omitempty"`

	// 网络研讨会实际召开时间（UTC时间），格式“yyyy-MM-dd HH:mm”。
	ActualStartTime *string `json:"actualStartTime,omitempty"`

	// 网络研讨会结束时间（UTC时间），格式“yyyy-MM-dd HH:mm”。
	EndTime *string `json:"endTime,omitempty"`

	// 网络研讨会实际时长，单位分钟。
	ActualDuration *int32 `json:"actualDuration,omitempty"`

	// 与会人数。
	AttendeeCount *int32 `json:"attendeeCount,omitempty"`

	// 主持人数。
	ChairCount *int32 `json:"chairCount,omitempty"`

	// 嘉宾数。
	GuestCount *int32 `json:"guestCount,omitempty"`

	// 观众人数。
	AudienceCount *int32 `json:"audienceCount,omitempty"`

	// VMR ID。
	VmrId *string `json:"vmrId,omitempty"`

	// 网络研讨会VMR最大观众数。
	VmrPkgAudienceParties *int32 `json:"vmrPkgAudienceParties,omitempty"`

	// 网络研讨会VMR名称。
	VmrPkgName *string `json:"vmrPkgName,omitempty"`
}

func (o OpenWebinarHistoryInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpenWebinarHistoryInfo struct{}"
	}

	return strings.Join([]string{"OpenWebinarHistoryInfo", string(data)}, " ")
}
