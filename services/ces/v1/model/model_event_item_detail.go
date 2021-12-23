package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

//
type EventItemDetail struct {
	// 事件内容，最大长度4096。

	Content *string `json:"content,omitempty"`
	// 所属分组。  资源分组对应的ID，必须传存在的分组ID。

	GroupId *string `json:"group_id,omitempty"`
	// 资源ID，支持字母、数字_ -：，最大长度128。

	ResourceId *string `json:"resource_id,omitempty"`
	// 资源名称，支持字母 中文 数字_ -. ，最大长度128。

	ResourceName *string `json:"resource_name,omitempty"`
	// 事件状态。  枚举类型：normal\\warning\\incident

	EventState *EventItemDetailEventState `json:"event_state,omitempty"`
	// 事件级别。  枚举类型：Critical, Major, Minor, Info

	EventLevel *EventItemDetailEventLevel `json:"event_level,omitempty"`
	// 事件用户。  支持字母 数字_ -/空格 ，最大长度64。

	EventUser *string `json:"event_user,omitempty"`
}

func (o EventItemDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventItemDetail struct{}"
	}

	return strings.Join([]string{"EventItemDetail", string(data)}, " ")
}

type EventItemDetailEventState struct {
	value string
}

type EventItemDetailEventStateEnum struct {
	NORMAL   EventItemDetailEventState
	WARNING  EventItemDetailEventState
	INCIDENT EventItemDetailEventState
}

func GetEventItemDetailEventStateEnum() EventItemDetailEventStateEnum {
	return EventItemDetailEventStateEnum{
		NORMAL: EventItemDetailEventState{
			value: "normal",
		},
		WARNING: EventItemDetailEventState{
			value: "warning",
		},
		INCIDENT: EventItemDetailEventState{
			value: "incident",
		},
	}
}

func (c EventItemDetailEventState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EventItemDetailEventState) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type EventItemDetailEventLevel struct {
	value string
}

type EventItemDetailEventLevelEnum struct {
	CRITICAL EventItemDetailEventLevel
	MAJOR    EventItemDetailEventLevel
	MINOR    EventItemDetailEventLevel
	INFO     EventItemDetailEventLevel
}

func GetEventItemDetailEventLevelEnum() EventItemDetailEventLevelEnum {
	return EventItemDetailEventLevelEnum{
		CRITICAL: EventItemDetailEventLevel{
			value: "Critical",
		},
		MAJOR: EventItemDetailEventLevel{
			value: "Major",
		},
		MINOR: EventItemDetailEventLevel{
			value: "Minor",
		},
		INFO: EventItemDetailEventLevel{
			value: "Info",
		},
	}
}

func (c EventItemDetailEventLevel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EventItemDetailEventLevel) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
