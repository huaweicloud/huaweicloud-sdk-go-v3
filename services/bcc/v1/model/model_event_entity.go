package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EventEntity 事件实体类
type EventEntity struct {

	// 事件ID
	EventId string `json:"event_id"`

	// 事件名称
	EventName *string `json:"event_name,omitempty"`

	// 事件类型
	EventType *string `json:"event_type,omitempty"`

	EventSource *EventSourceEnum `json:"event_source,omitempty"`

	// 事件产生时间
	Time *string `json:"time,omitempty"`

	Detail *EventDetailEntity `json:"detail,omitempty"`
}

func (o EventEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventEntity struct{}"
	}

	return strings.Join([]string{"EventEntity", string(data)}, " ")
}
