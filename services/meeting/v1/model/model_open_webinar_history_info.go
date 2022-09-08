package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 历史召开研讨会议信息
type OpenWebinarHistoryInfo struct {

	// 会议id
	ConferenceId *string `json:"conferenceId,omitempty"`

	// 会议UUID
	ConfUUID *string `json:"confUUID,omitempty"`

	// 主题
	Subject *string `json:"subject,omitempty"`

	// 会议订阅者
	ScheduserName *string `json:"scheduserName,omitempty"`

	// 会议主持人
	Moderator *string `json:"moderator,omitempty"`

	// 部门名称
	DeptName *string `json:"deptName,omitempty"`

	// 时区ID
	TimeZoneId *int32 `json:"timeZoneId,omitempty"`

	// 会议预约时间
	StartTime *string `json:"startTime,omitempty"`

	// 预约会议时长，单位分钟
	Duration *int32 `json:"duration,omitempty"`

	// 会议开始时间
	ActualStartTime *string `json:"actualStartTime,omitempty"`

	// 会议结束时间
	EndTime *string `json:"endTime,omitempty"`

	// 实际会议时长，单位分钟
	ActualDuration *int32 `json:"actualDuration,omitempty"`

	// 与会人数
	AttendeeCount *int32 `json:"attendeeCount,omitempty"`

	// 主持人人数
	ChairCount *int32 `json:"chairCount,omitempty"`

	// 嘉宾人数
	GuestCount *int32 `json:"guestCount,omitempty"`

	// 观众人数
	AudienceCount *int32 `json:"audienceCount,omitempty"`

	// VMR ID
	VmrId *string `json:"vmrId,omitempty"`

	// VMR资源规格-最大观众数
	VmrPkgAudienceParties *int32 `json:"vmrPkgAudienceParties,omitempty"`

	// 网络研讨会资源名
	VmrPkgName *string `json:"vmrPkgName,omitempty"`
}

func (o OpenWebinarHistoryInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpenWebinarHistoryInfo struct{}"
	}

	return strings.Join([]string{"OpenWebinarHistoryInfo", string(data)}, " ")
}
