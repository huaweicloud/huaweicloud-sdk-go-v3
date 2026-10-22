package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UrlMatchDto URL 匹配规则
type UrlMatchDto struct {

	// 匹配算子（EQUAL/CONTAIN）
	Operator UrlMatchDtoOperator `json:"operator"`

	// 匹配值（URL 或 URL 片段）
	Value string `json:"value"`
}

func (o UrlMatchDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UrlMatchDto struct{}"
	}

	return strings.Join([]string{"UrlMatchDto", string(data)}, " ")
}

type UrlMatchDtoOperator struct {
	value string
}

type UrlMatchDtoOperatorEnum struct {
	EQUAL   UrlMatchDtoOperator
	CONTAIN UrlMatchDtoOperator
}

func GetUrlMatchDtoOperatorEnum() UrlMatchDtoOperatorEnum {
	return UrlMatchDtoOperatorEnum{
		EQUAL: UrlMatchDtoOperator{
			value: "EQUAL",
		},
		CONTAIN: UrlMatchDtoOperator{
			value: "CONTAIN",
		},
	}
}

func (c UrlMatchDtoOperator) Value() string {
	return c.value
}

func (c UrlMatchDtoOperator) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UrlMatchDtoOperator) UnmarshalJSON(b []byte) error {
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
