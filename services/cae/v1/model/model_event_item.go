package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type EventItem struct {

	// 事件名称。
	Name *string `json:"name,omitempty"`

	// 涉及对象类型。
	InvolvedObjectKind *EventItemInvolvedObjectKind `json:"involved_object_kind,omitempty"`

	// 涉及对象。
	InvolvedObject *string `json:"involved_object,omitempty"`

	// 组件事件信息。
	Message *string `json:"message,omitempty"`

	// 创建时间。
	CreatedAt *string `json:"created_at,omitempty"`

	// 更新时间。
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

type EventItemInvolvedObjectKind struct {
	value string
}

type EventItemInvolvedObjectKindEnum struct {
	COMPONENT          EventItemInvolvedObjectKind
	COMPONENT_INSTANCE EventItemInvolvedObjectKind
	COMPONENT_SCALING  EventItemInvolvedObjectKind
}

func GetEventItemInvolvedObjectKindEnum() EventItemInvolvedObjectKindEnum {
	return EventItemInvolvedObjectKindEnum{
		COMPONENT: EventItemInvolvedObjectKind{
			value: "Component",
		},
		COMPONENT_INSTANCE: EventItemInvolvedObjectKind{
			value: "ComponentInstance",
		},
		COMPONENT_SCALING: EventItemInvolvedObjectKind{
			value: "ComponentScaling",
		},
	}
}

func (c EventItemInvolvedObjectKind) Value() string {
	return c.value
}

func (c EventItemInvolvedObjectKind) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EventItemInvolvedObjectKind) UnmarshalJSON(b []byte) error {
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
