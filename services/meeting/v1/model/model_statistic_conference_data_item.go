package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 会议总体数据的单个时间点数据。
type StatisticConferenceDataItem struct {

	// 日期/月份，category = conference_info时有效。 小时，category = conference_hourly_info时有效。
	Time *string `json:"time,omitempty"`

	// 会议数(含VMR)。 category = conference_info时有效。
	ConfCount *string `json:"confCount,omitempty"`

	// 会议时长(秒)(含VMR)。 category = conference_info时有效。
	ConfDuration *string `json:"confDuration,omitempty"`

	// 与会人次(含VMR)。 category = conference_info时有效。
	AttendeeCount *string `json:"attendeeCount,omitempty"`

	// 并发会议使用数。 category = conference_info时有效。
	ConfConcurrentUsedCount *string `json:"confConcurrentUsedCount,omitempty"`

	// 小时单位会议数(含VMR)。 category = conference_hourly_info时有效。
	Conf24hCount *string `json:"conf24hCount,omitempty"`

	// 小时单位与会人次(含VMR)。 category = conference_hourly_info时有效。
	Conf24hAttendeeCount *string `json:"conf24hAttendeeCount,omitempty"`
}

func (o StatisticConferenceDataItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StatisticConferenceDataItem struct{}"
	}

	return strings.Join([]string{"StatisticConferenceDataItem", string(data)}, " ")
}
