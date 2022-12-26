package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 定时任务数据
type Trigger struct {

	// 定时策略。once、corn、periodic
	Policy string `json:"policy"`

	// 触发器执行时间。单次执行为UTC毫秒数、简单周期为\"[\\\"7\\\"]\"、corn为corn表达式\"0 23 * * *\"
	ScheduledTime string `json:"scheduled_time"`

	// 时区。
	TimeZone string `json:"time_zone"`

	// smn主题urn。
	TopicUrn string `json:"topic_urn"`
}

func (o Trigger) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Trigger struct{}"
	}

	return strings.Join([]string{"Trigger", string(data)}, " ")
}
