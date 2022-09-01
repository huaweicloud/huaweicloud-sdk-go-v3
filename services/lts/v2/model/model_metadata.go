package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Metadata struct {

	// 告警类型
	EventType string `json:"event_type" xml:"event_type"`

	// 告警id
	EventId string `json:"event_id" xml:"event_id"`

	// 告警级别
	EventSeverity string `json:"event_severity" xml:"event_severity"`

	// 告警名称
	EventName string `json:"event_name" xml:"event_name"`

	// 资源类型
	ResourceType string `json:"resource_type" xml:"resource_type"`

	// 日志组/流名称
	ResourceId string `json:"resource_id" xml:"resource_id"`

	// 告警源
	ResourceProvider string `json:"resource_provider" xml:"resource_provider"`

	// 告警规则类型(SQL/关键词)
	LtsAlarmType string `json:"lts_alarm_type" xml:"lts_alarm_type"`
}

func (o Metadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Metadata struct{}"
	}

	return strings.Join([]string{"Metadata", string(data)}, " ")
}
