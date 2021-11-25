package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 指标具体数值。
type ValueData struct {
	// 指标名称。

	MetricName string `json:"metric_name"`
	// 取值范围 只能是\"int\"或\"float\"。 数据的类型。

	Type *ValueDataType `json:"type,omitempty"`
	// 数据的单位。

	Unit *string `json:"unit,omitempty"`
	// 取值范围 有效的数值类型 指标数据的值。

	Value float64 `json:"value"`
}

func (o ValueData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValueData struct{}"
	}

	return strings.Join([]string{"ValueData", string(data)}, " ")
}

type ValueDataType struct {
	value string
}

type ValueDataTypeEnum struct {
	INT   ValueDataType
	FLOAT ValueDataType
}

func GetValueDataTypeEnum() ValueDataTypeEnum {
	return ValueDataTypeEnum{
		INT: ValueDataType{
			value: "int",
		},
		FLOAT: ValueDataType{
			value: "float",
		},
	}
}

func (c ValueDataType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ValueDataType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
