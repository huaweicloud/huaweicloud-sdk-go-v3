package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Event 事件通知对象。
type Event struct {

	// 事件通知名称。
	Name *string `json:"name,omitempty"`

	// 事件通知的资源标识符。
	EventId *string `json:"event_id,omitempty"`

	// 设置事件的基础事件类型列表,。  约束：数组大小：最小1，最大12。
	EventTypes *[]EventEventTypes `json:"event_types,omitempty"`

	// 事件通知状态，取值如下。  ENABLED：表示启用状态 DISABLED：表示禁用状态
	State *EventState `json:"state,omitempty"`

	// 事件通知创建时间，时间戳，即从1970年1月1日至该时间的总秒数。
	CreateTime *int64 `json:"create_time,omitempty"`

	// 事件通知上次更新时间，时间戳，即从1970年1月1日至该时间的总秒数。
	UpdateTime *int64 `json:"update_time,omitempty"`

	Notification *Notification `json:"notification,omitempty"`
}

func (o Event) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Event struct{}"
	}

	return strings.Join([]string{"Event", string(data)}, " ")
}

type EventEventTypes struct {
	value string
}

type EventEventTypesEnum struct {
	SECRET_VERSION_CREATED EventEventTypes
	SECRET_VERSION_EXPIRED EventEventTypes
	SECRET_ROTATED         EventEventTypes
	SECRET_DELETED         EventEventTypes
	SECRET_ROTATED_FAILED  EventEventTypes
}

func GetEventEventTypesEnum() EventEventTypesEnum {
	return EventEventTypesEnum{
		SECRET_VERSION_CREATED: EventEventTypes{
			value: "SECRET_VERSION_CREATED",
		},
		SECRET_VERSION_EXPIRED: EventEventTypes{
			value: "SECRET_VERSION_EXPIRED",
		},
		SECRET_ROTATED: EventEventTypes{
			value: "SECRET_ROTATED",
		},
		SECRET_DELETED: EventEventTypes{
			value: "SECRET_DELETED",
		},
		SECRET_ROTATED_FAILED: EventEventTypes{
			value: "SECRET_ROTATED_FAILED",
		},
	}
}

func (c EventEventTypes) Value() string {
	return c.value
}

func (c EventEventTypes) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EventEventTypes) UnmarshalJSON(b []byte) error {
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

type EventState struct {
	value string
}

type EventStateEnum struct {
	ENABLED  EventState
	DISABLED EventState
}

func GetEventStateEnum() EventStateEnum {
	return EventStateEnum{
		ENABLED: EventState{
			value: "ENABLED",
		},
		DISABLED: EventState{
			value: "DISABLED",
		},
	}
}

func (c EventState) Value() string {
	return c.value
}

func (c EventState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EventState) UnmarshalJSON(b []byte) error {
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
