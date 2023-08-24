package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type VariableAttributes struct {

	// 参数索引。
	VariableIndex int32 `json:"variable_index"`

	// 参数类型。
	VariableType VariableAttributesVariableType `json:"variable_type"`

	// 参数描述。变量类型为TEXT（其他）时必填。
	VariableDesc *string `json:"variable_desc,omitempty"`
}

func (o VariableAttributes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VariableAttributes struct{}"
	}

	return strings.Join([]string{"VariableAttributes", string(data)}, " ")
}

type VariableAttributesVariableType struct {
	value string
}

type VariableAttributesVariableTypeEnum struct {
	PHONE     VariableAttributesVariableType
	NEWTEXT   VariableAttributesVariableType
	CHARDIGIT VariableAttributesVariableType
	DATETIME  VariableAttributesVariableType
	MONEY     VariableAttributesVariableType
	TEXT      VariableAttributesVariableType
}

func GetVariableAttributesVariableTypeEnum() VariableAttributesVariableTypeEnum {
	return VariableAttributesVariableTypeEnum{
		PHONE: VariableAttributesVariableType{
			value: "PHONE：电话号码",
		},
		NEWTEXT: VariableAttributesVariableType{
			value: "NEWTEXT：解析标识",
		},
		CHARDIGIT: VariableAttributesVariableType{
			value: "CHARDIGIT：其他号码(如验证码、订单号、密码等)",
		},
		DATETIME: VariableAttributesVariableType{
			value: "DATETIME：日期时间",
		},
		MONEY: VariableAttributesVariableType{
			value: "MONEY：金额",
		},
		TEXT: VariableAttributesVariableType{
			value: "TEXT：其他",
		},
	}
}

func (c VariableAttributesVariableType) Value() string {
	return c.value
}

func (c VariableAttributesVariableType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VariableAttributesVariableType) UnmarshalJSON(b []byte) error {
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
