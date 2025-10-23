package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// EventStateEnum 事件状态，取值范围：normal,warning,incident
type EventStateEnum struct {
	value string
}

type EventStateEnumEnum struct {
	NORMAL   EventStateEnum
	WARNING  EventStateEnum
	INCIDENT EventStateEnum
}

func GetEventStateEnumEnum() EventStateEnumEnum {
	return EventStateEnumEnum{
		NORMAL: EventStateEnum{
			value: "normal",
		},
		WARNING: EventStateEnum{
			value: "warning",
		},
		INCIDENT: EventStateEnum{
			value: "incident",
		},
	}
}

func (c EventStateEnum) Value() string {
	return c.value
}

func (c EventStateEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EventStateEnum) UnmarshalJSON(b []byte) error {
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
