package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SpecType 资源规格类型： 1.APU：APU类型, 2.DPU：DPU类型, 3.MU：MU类型,  4.RESOURCE_USAGE_UNIT：资源使用量类型,  5.TOKENS_USAGE_UNIT：Tokens数使用量类型
type SpecType struct {
	value string
}

type SpecTypeEnum struct {
	APU                 SpecType
	DPU                 SpecType
	MU                  SpecType
	RESOURCE_USAGE_UNIT SpecType
	TOKENS_USAGE_UNIT   SpecType
}

func GetSpecTypeEnum() SpecTypeEnum {
	return SpecTypeEnum{
		APU: SpecType{
			value: "APU",
		},
		DPU: SpecType{
			value: "DPU",
		},
		MU: SpecType{
			value: "MU",
		},
		RESOURCE_USAGE_UNIT: SpecType{
			value: "RESOURCE_USAGE_UNIT",
		},
		TOKENS_USAGE_UNIT: SpecType{
			value: "TOKENS_USAGE_UNIT",
		},
	}
}

func (c SpecType) Value() string {
	return c.value
}

func (c SpecType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SpecType) UnmarshalJSON(b []byte) error {
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
