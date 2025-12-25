package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// IsapTimeUnit **参数解释**: 时间单位 - MINUTE 分钟 - HOUR 小时 - DAY 天 - MONTH 月  **约束限制** 不涉及 **取值范围**: - MINUTE - HOUR - DAY - MONTH  **默认值** 不涉及
type IsapTimeUnit struct {
	value string
}

type IsapTimeUnitEnum struct {
	MINUTE IsapTimeUnit
	HOUR   IsapTimeUnit
	DAY    IsapTimeUnit
	MONTH  IsapTimeUnit
}

func GetIsapTimeUnitEnum() IsapTimeUnitEnum {
	return IsapTimeUnitEnum{
		MINUTE: IsapTimeUnit{
			value: "MINUTE",
		},
		HOUR: IsapTimeUnit{
			value: "HOUR",
		},
		DAY: IsapTimeUnit{
			value: "DAY",
		},
		MONTH: IsapTimeUnit{
			value: "MONTH",
		},
	}
}

func (c IsapTimeUnit) Value() string {
	return c.value
}

func (c IsapTimeUnit) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *IsapTimeUnit) UnmarshalJSON(b []byte) error {
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
