package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type EventItem struct {

	// 事件名。
	Name *string `json:"name,omitempty"`

	// 涉及对象。
	InvolvedObject *EventItemInvolvedObject `json:"involved_object,omitempty"`

	// 组件事件信息。
	Message *string `json:"message,omitempty"`

	// 创建时间。
	CreatedAt *string `json:"created_at,omitempty"`

	// 修改时间。
	UpdatedAt *string `json:"updated_at,omitempty"`

	// 组件事件状态。
	Status *EventItemStatus `json:"status,omitempty"`

	// 事件发生次数。
	Count *int32 `json:"count,omitempty"`
}

func (o EventItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventItem struct{}"
	}

	return strings.Join([]string{"EventItem", string(data)}, " ")
}

type EventItemInvolvedObject struct {
	value string
}

type EventItemInvolvedObjectEnum struct {
	COMPONENT          EventItemInvolvedObject
	COMPONENT_INSTANCE EventItemInvolvedObject
	COMPONENT_SCALING  EventItemInvolvedObject
}

func GetEventItemInvolvedObjectEnum() EventItemInvolvedObjectEnum {
	return EventItemInvolvedObjectEnum{
		COMPONENT: EventItemInvolvedObject{
			value: "Component",
		},
		COMPONENT_INSTANCE: EventItemInvolvedObject{
			value: "ComponentInstance",
		},
		COMPONENT_SCALING: EventItemInvolvedObject{
			value: "ComponentScaling",
		},
	}
}

func (c EventItemInvolvedObject) Value() string {
	return c.value
}

func (c EventItemInvolvedObject) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EventItemInvolvedObject) UnmarshalJSON(b []byte) error {
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

type EventItemStatus struct {
	value string
}

type EventItemStatusEnum struct {
	WARNING EventItemStatus
	NORMAL  EventItemStatus
}

func GetEventItemStatusEnum() EventItemStatusEnum {
	return EventItemStatusEnum{
		WARNING: EventItemStatus{
			value: "Warning",
		},
		NORMAL: EventItemStatus{
			value: "Normal",
		},
	}
}

func (c EventItemStatus) Value() string {
	return c.value
}

func (c EventItemStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EventItemStatus) UnmarshalJSON(b []byte) error {
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
