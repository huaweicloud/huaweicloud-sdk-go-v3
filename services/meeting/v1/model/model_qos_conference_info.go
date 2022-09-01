package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QoS会议信息。
type QosConferenceInfo struct {

	// 会议UUID。
	ConfUUID *string `json:"confUUID,omitempty" xml:"confUUID"`

	// 会议ID。
	ConferenceID *string `json:"conferenceID,omitempty" xml:"conferenceID"`

	// 会议主题。
	Subject *string `json:"subject,omitempty" xml:"subject"`

	// 会议预订者名称。
	ScheduserName *string `json:"scheduserName,omitempty" xml:"scheduserName"`

	// 部门。
	DeptName *string `json:"deptName,omitempty" xml:"deptName"`

	// 总体告警 YES/NO。 说明： * 会议的音频，视频，屏幕共享，CPU任一项产生告警，总体告警就为YES。
	Alarm *string `json:"alarm,omitempty" xml:"alarm"`

	// 音频告警 YES/NO。 说明： * 会议中任一与会者存在音频告警，会议音频告警就为YES。
	AudioAlarm *string `json:"audioAlarm,omitempty" xml:"audioAlarm"`

	// 视频告警 YES/NO。 说明： * 会议中任一与会者存在视频告警，会议视频告警就为YES。
	VideoAlarm *string `json:"videoAlarm,omitempty" xml:"videoAlarm"`

	// 屏幕共享告警 YES/NO。 说明： * 会议中任一与会者存在屏幕共享告警，会议屏幕共享告警就为YES。
	ScreenAlarm *string `json:"screenAlarm,omitempty" xml:"screenAlarm"`

	// CPU告警 YES/NO。 说明： * 会议中任一与会者存在CPU告警，会议CPU告警就为YES。'
	CpuAlarm *string `json:"cpuAlarm,omitempty" xml:"cpuAlarm"`

	// 时区。详情参考时区表（云会议帮助中心->服务端API参考->附录->时区表），中国默认时区56-东八区。
	TimeZoneID *string `json:"timeZoneID,omitempty" xml:"timeZoneID"`

	// 会议开始时间(UTC时间), Unix时间戳（单位毫秒）。
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime"`

	// 会议结束时间(UTC时间), Unix时间戳（单位毫秒）。 说明： * 在线会议：会议召开中，endTime = 会议预计结束时间。 * 历史会议：会议已结束，endTime = 会议实际结束时间。
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime"`

	// 会议召开时长（分钟）。 说明： * 在线会议：会议召开中，duration = 0。 * 历史会议：会议已结束，duration = 会议实际召开时间。
	Duration *int32 `json:"duration,omitempty" xml:"duration"`

	// 与会方数。 说明： * 同一用户多次进出会议属于不同的与会，与会方数计算多次。
	Participants *int32 `json:"participants,omitempty" xml:"participants"`

	// 是否网络研讨会。
	Webinar *bool `json:"webinar,omitempty" xml:"webinar"`
}

func (o QosConferenceInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QosConferenceInfo struct{}"
	}

	return strings.Join([]string{"QosConferenceInfo", string(data)}, " ")
}
