package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type EventInfo struct {

	// 事件ID
	EventId *string `json:"event_id,omitempty"`

	// 描述ID
	MessageId *string `json:"message_id,omitempty"`

	// 事件描述
	EventMessage *string `json:"event_message,omitempty"`

	// 事件级别，枚举类型：critical, major, minor, info
	EventLevel *EventInfoEventLevel `json:"event_level,omitempty"`

	// 事件状态
	Status *string `json:"status,omitempty"`

	// 资源ID，最大长度128
	ResourceId *string `json:"resource_id,omitempty"`

	// 事件发生时间，UNIX时间戳，单位毫秒
	Time *int64 `json:"time,omitempty"`

	// 原始事件，最大长度4096
	OriginEvent *string `json:"origin_event,omitempty"`
}

func (o EventInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventInfo struct{}"
	}

	return strings.Join([]string{"EventInfo", string(data)}, " ")
}

type EventInfoEventLevel struct {
	value string
}

type EventInfoEventLevelEnum struct {
	CRITICAL EventInfoEventLevel
	MAJOR    EventInfoEventLevel
	MINOR    EventInfoEventLevel
	INFO     EventInfoEventLevel
}

func GetEventInfoEventLevelEnum() EventInfoEventLevelEnum {
	return EventInfoEventLevelEnum{
		CRITICAL: EventInfoEventLevel{
			value: "critical",
		},
		MAJOR: EventInfoEventLevel{
			value: "major",
		},
		MINOR: EventInfoEventLevel{
			value: "minor",
		},
		INFO: EventInfoEventLevel{
			value: "info",
		},
	}
}

func (c EventInfoEventLevel) Value() string {
	return c.value
}

func (c EventInfoEventLevel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EventInfoEventLevel) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
