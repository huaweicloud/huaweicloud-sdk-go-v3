package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ScheduledTypeEnum 执行周期： * `FIXED_TIME` - 指定时间 * `DAY` - 按天 * `WEEK` - 按周 * `MONTH` - 按月
type ScheduledTypeEnum struct {
	value string
}

type ScheduledTypeEnumEnum struct {
	FIXED_TIME ScheduledTypeEnum
	DAY        ScheduledTypeEnum
	WEEK       ScheduledTypeEnum
	MONTH      ScheduledTypeEnum
}

func GetScheduledTypeEnumEnum() ScheduledTypeEnumEnum {
	return ScheduledTypeEnumEnum{
		FIXED_TIME: ScheduledTypeEnum{
			value: "FIXED_TIME",
		},
		DAY: ScheduledTypeEnum{
			value: "DAY",
		},
		WEEK: ScheduledTypeEnum{
			value: "WEEK",
		},
		MONTH: ScheduledTypeEnum{
			value: "MONTH",
		},
	}
}

func (c ScheduledTypeEnum) Value() string {
	return c.value
}

func (c ScheduledTypeEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ScheduledTypeEnum) UnmarshalJSON(b []byte) error {
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
