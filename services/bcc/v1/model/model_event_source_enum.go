package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// EventSourceEnum 事件来源，取值范围：SYS.CBR,SYS.RDS,SYS.GaussDB
type EventSourceEnum struct {
	value string
}

type EventSourceEnumEnum struct {
	SYS_CBR      EventSourceEnum
	SYS_RDS      EventSourceEnum
	SYS_GAUSS_DB EventSourceEnum
}

func GetEventSourceEnumEnum() EventSourceEnumEnum {
	return EventSourceEnumEnum{
		SYS_CBR: EventSourceEnum{
			value: "SYS.CBR",
		},
		SYS_RDS: EventSourceEnum{
			value: "SYS.RDS",
		},
		SYS_GAUSS_DB: EventSourceEnum{
			value: "SYS.GaussDB",
		},
	}
}

func (c EventSourceEnum) Value() string {
	return c.value
}

func (c EventSourceEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EventSourceEnum) UnmarshalJSON(b []byte) error {
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
