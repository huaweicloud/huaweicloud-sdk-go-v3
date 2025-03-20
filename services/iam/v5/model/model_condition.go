package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Condition 条件键。
type Condition struct {

	// 条件键的名称。
	Key string `json:"key"`

	// 条件值的数据类型。
	ValueType ConditionValueType `json:"value_type"`

	// 条件值是否为多值。
	MultiValued bool `json:"multi_valued"`

	Description *Description `json:"description"`
}

func (o Condition) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Condition struct{}"
	}

	return strings.Join([]string{"Condition", string(data)}, " ")
}

type ConditionValueType struct {
	value string
}

type ConditionValueTypeEnum struct {
	STRING     ConditionValueType
	NUMERIC    ConditionValueType
	DATE       ConditionValueType
	BOOLEAN    ConditionValueType
	IP_ADDRESS ConditionValueType
}

func GetConditionValueTypeEnum() ConditionValueTypeEnum {
	return ConditionValueTypeEnum{
		STRING: ConditionValueType{
			value: "string",
		},
		NUMERIC: ConditionValueType{
			value: "numeric",
		},
		DATE: ConditionValueType{
			value: "date",
		},
		BOOLEAN: ConditionValueType{
			value: "boolean",
		},
		IP_ADDRESS: ConditionValueType{
			value: "ip_address",
		},
	}
}

func (c ConditionValueType) Value() string {
	return c.value
}

func (c ConditionValueType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ConditionValueType) UnmarshalJSON(b []byte) error {
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
