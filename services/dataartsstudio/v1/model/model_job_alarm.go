package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// JobAlarm 作业告警信息
type JobAlarm struct {

	// 告警类型。
	AlarmType *string `json:"alarmType,omitempty"`

	// 消息通知主题URN。
	TopicUrn *string `json:"topicUrn,omitempty"`
}

func (o JobAlarm) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobAlarm struct{}"
	}

	return strings.Join([]string{"JobAlarm", string(data)}, " ")
}
