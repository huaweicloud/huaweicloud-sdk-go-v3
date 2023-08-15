package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// TemplateType 模板类型(custom代表默认自定义模板，system代表系统模板)
type TemplateType struct {
	value string
}

type TemplateTypeEnum struct {
	SYSTEM TemplateType
	CUSTOM TemplateType
}

func GetTemplateTypeEnum() TemplateTypeEnum {
	return TemplateTypeEnum{
		SYSTEM: TemplateType{
			value: "system",
		},
		CUSTOM: TemplateType{
			value: "custom",
		},
	}
}

func (c TemplateType) Value() string {
	return c.value
}

func (c TemplateType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TemplateType) UnmarshalJSON(b []byte) error {
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
