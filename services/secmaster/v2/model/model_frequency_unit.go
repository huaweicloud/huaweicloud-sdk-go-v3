package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// FrequencyUnit **参数解释**: 频率单位 - MINUTE 分钟 - HOUR 小时 - DAY 天  **约束限制** 不涉及 **取值范围**: - MINUTE - HOUR - DAY  **默认值** 不涉及
type FrequencyUnit struct {
	value string
}

type FrequencyUnitEnum struct {
	MINUTE FrequencyUnit
	HOUR   FrequencyUnit
	DAY    FrequencyUnit
}

func GetFrequencyUnitEnum() FrequencyUnitEnum {
	return FrequencyUnitEnum{
		MINUTE: FrequencyUnit{
			value: "MINUTE",
		},
		HOUR: FrequencyUnit{
			value: "HOUR",
		},
		DAY: FrequencyUnit{
			value: "DAY",
		},
	}
}

func (c FrequencyUnit) Value() string {
	return c.value
}

func (c FrequencyUnit) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *FrequencyUnit) UnmarshalJSON(b []byte) error {
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
