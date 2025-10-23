package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AlarmTypeEnum 告警类型
type AlarmTypeEnum struct {
	value string
}

type AlarmTypeEnumEnum struct {
	EVENT_SYS      AlarmTypeEnum
	EVENT_CUSTOM   AlarmTypeEnum
	ALL_INSTANCE   AlarmTypeEnum
	MULTI_INSTANCE AlarmTypeEnum
	RESOURCE_GROUP AlarmTypeEnum
}

func GetAlarmTypeEnumEnum() AlarmTypeEnumEnum {
	return AlarmTypeEnumEnum{
		EVENT_SYS: AlarmTypeEnum{
			value: "EVENT.SYS",
		},
		EVENT_CUSTOM: AlarmTypeEnum{
			value: "EVENT.CUSTOM",
		},
		ALL_INSTANCE: AlarmTypeEnum{
			value: "ALL_INSTANCE",
		},
		MULTI_INSTANCE: AlarmTypeEnum{
			value: "MULTI_INSTANCE",
		},
		RESOURCE_GROUP: AlarmTypeEnum{
			value: "RESOURCE_GROUP",
		},
	}
}

func (c AlarmTypeEnum) Value() string {
	return c.value
}

func (c AlarmTypeEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AlarmTypeEnum) UnmarshalJSON(b []byte) error {
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
