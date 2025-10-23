package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// EventLevelEnum 事件级别，取值范围：Critical,Major,Minor,Info
type EventLevelEnum struct {
	value string
}

type EventLevelEnumEnum struct {
	CRITICAL EventLevelEnum
	MAJOR    EventLevelEnum
	MINOR    EventLevelEnum
	INFO     EventLevelEnum
}

func GetEventLevelEnumEnum() EventLevelEnumEnum {
	return EventLevelEnumEnum{
		CRITICAL: EventLevelEnum{
			value: "Critical",
		},
		MAJOR: EventLevelEnum{
			value: "Major",
		},
		MINOR: EventLevelEnum{
			value: "Minor",
		},
		INFO: EventLevelEnum{
			value: "Info",
		},
	}
}

func (c EventLevelEnum) Value() string {
	return c.value
}

func (c EventLevelEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EventLevelEnum) UnmarshalJSON(b []byte) error {
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
