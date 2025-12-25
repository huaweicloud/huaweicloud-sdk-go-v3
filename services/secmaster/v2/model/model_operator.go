package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Operator **参数解释**: 操作符类型 - GT 大于 - LT 小于 - EQ 等于 - NE 不等于  **约束限制** 不涉及 **取值范围**: - GT - LT - EQ - NE  **默认值** 不涉及
type Operator struct {
	value string
}

type OperatorEnum struct {
	GT Operator
	LT Operator
	EQ Operator
	NE Operator
}

func GetOperatorEnum() OperatorEnum {
	return OperatorEnum{
		GT: Operator{
			value: "GT",
		},
		LT: Operator{
			value: "LT",
		},
		EQ: Operator{
			value: "EQ",
		},
		NE: Operator{
			value: "NE",
		},
	}
}

func (c Operator) Value() string {
	return c.value
}

func (c Operator) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *Operator) UnmarshalJSON(b []byte) error {
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
