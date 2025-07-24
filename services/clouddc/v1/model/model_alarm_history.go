package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AlarmHistory struct {

	// 告警记录ID
	AlarmRecordId *string `json:"alarm_record_id,omitempty"`

	// 告警ID
	AlarmId *string `json:"alarm_id,omitempty"`

	// 告警名称
	AlarmName *string `json:"alarm_name,omitempty"`

	// 告警状态
	AlarmStatus *string `json:"alarm_status,omitempty"`

	// 告警级别
	AlarmLevel *int32 `json:"alarm_level,omitempty"`

	// 告警开始时间，Unix时间戳
	BeginTime *int64 `json:"begin_time,omitempty"`

	// 告警结束时间，Unix时间戳
	EndTime *int64 `json:"end_time,omitempty"`

	// 最后一次告警时间，Unix时间戳
	LastAlarmTime *int64 `json:"last_alarm_time,omitempty"`

	Metric *Metric `json:"metric,omitempty"`
}

func (o AlarmHistory) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmHistory struct{}"
	}

	return strings.Join([]string{"AlarmHistory", string(data)}, " ")
}
