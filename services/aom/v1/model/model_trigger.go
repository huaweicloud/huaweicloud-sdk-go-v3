package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 定时任务数据
type Trigger struct {

	// 定时任务的必填信息。
	Information string `json:"information"`

	// 触发器执行时间。
	ScheduledTime string `json:"scheduled_time"`

	// 时区。
	TimeZone string `json:"time_zone"`

	// smn主题urn。
	TopicUrn string `json:"topic_urn"`

	// aom告警名字。
	AlarmName string `json:"alarm_name"`
}

func (o Trigger) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Trigger struct{}"
	}

	return strings.Join([]string{"Trigger", string(data)}, " ")
}
