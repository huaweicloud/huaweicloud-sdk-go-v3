package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AlarmStatusEnum 告警状态，取值范围：ok,alarm,invalid
type AlarmStatusEnum struct {
	value string
}

type AlarmStatusEnumEnum struct {
	OK      AlarmStatusEnum
	ALARM   AlarmStatusEnum
	INVALID AlarmStatusEnum
}

func GetAlarmStatusEnumEnum() AlarmStatusEnumEnum {
	return AlarmStatusEnumEnum{
		OK: AlarmStatusEnum{
			value: "ok",
		},
		ALARM: AlarmStatusEnum{
			value: "alarm",
		},
		INVALID: AlarmStatusEnum{
			value: "invalid",
		},
	}
}

func (c AlarmStatusEnum) Value() string {
	return c.value
}

func (c AlarmStatusEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AlarmStatusEnum) UnmarshalJSON(b []byte) error {
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
