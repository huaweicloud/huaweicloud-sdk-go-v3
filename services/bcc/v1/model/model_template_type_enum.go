package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// TemplateTypeEnum 模板类型枚举
type TemplateTypeEnum struct {
	value string
}

type TemplateTypeEnumEnum struct {
	SYSTEM TemplateTypeEnum
	CUSTOM TemplateTypeEnum
}

func GetTemplateTypeEnumEnum() TemplateTypeEnumEnum {
	return TemplateTypeEnumEnum{
		SYSTEM: TemplateTypeEnum{
			value: "system",
		},
		CUSTOM: TemplateTypeEnum{
			value: "custom",
		},
	}
}

func (c TemplateTypeEnum) Value() string {
	return c.value
}

func (c TemplateTypeEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TemplateTypeEnum) UnmarshalJSON(b []byte) error {
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
