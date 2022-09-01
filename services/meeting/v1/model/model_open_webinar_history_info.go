package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 历史召开研讨会议信息
type OpenWebinarHistoryInfo struct {

	// 会议id
	ConferenceId *string `json:"conferenceId,omitempty" xml:"conferenceId"`

	// 会议UUID
	ConfUUID *string `json:"confUUID,omitempty" xml:"confUUID"`

	// 主题
	Subject *string `json:"subject,omitempty" xml:"subject"`

	// 会议订阅者
	ScheduserName *string `json:"scheduserName,omitempty" xml:"scheduserName"`

	// 会议主持人
	Moderator *string `json:"moderator,omitempty" xml:"moderator"`

	// 部门名称
	DeptName *string `json:"deptName,omitempty" xml:"deptName"`

	// 时区ID
	TimeZoneId *int32 `json:"timeZoneId,omitempty" xml:"timeZoneId"`

	// 会议预约时间
	StartTime *string `json:"startTime,omitempty" xml:"startTime"`

	// 预约会议时长，单位分钟
	Duration *int32 `json:"duration,omitempty" xml:"duration"`

	// 会议开始时间
	ActualStartTime *string `json:"actualStartTime,omitempty" xml:"actualStartTime"`

	// 会议结束时间
	EndTime *string `json:"endTime,omitempty" xml:"endTime"`

	// 实际会议时长，单位分钟
	ActualDuration *int32 `json:"actualDuration,omitempty" xml:"actualDuration"`

	// 与会人数
	AttendeeCount *int32 `json:"attendeeCount,omitempty" xml:"attendeeCount"`

	// 主持人人数
	ChairCount *int32 `json:"chairCount,omitempty" xml:"chairCount"`

	// 嘉宾人数
	GuestCount *int32 `json:"guestCount,omitempty" xml:"guestCount"`

	// 观众人数
	AudienceCount *int32 `json:"audienceCount,omitempty" xml:"audienceCount"`

	// VMR ID
	VmrId *string `json:"vmrId,omitempty" xml:"vmrId"`

	// VMR资源规格-最大观众数
	VmrPkgAudienceParties *int32 `json:"vmrPkgAudienceParties,omitempty" xml:"vmrPkgAudienceParties"`

	// 网络研讨会资源名
	VmrPkgName *string `json:"vmrPkgName,omitempty" xml:"vmrPkgName"`
}

func (o OpenWebinarHistoryInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpenWebinarHistoryInfo struct{}"
	}

	return strings.Join([]string{"OpenWebinarHistoryInfo", string(data)}, " ")
}
