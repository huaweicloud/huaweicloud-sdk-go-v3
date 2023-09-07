package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type JobAlarm struct {

	// 告警类型
	AlarmType string `json:"alarmType"`

	TopicUrn string `json:"topicUrn"`
}

func (o JobAlarm) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobAlarm struct{}"
	}

	return strings.Join([]string{"JobAlarm", string(data)}, " ")
}
