package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 模板类型(custom代表默认自定义模板，system代表系统模板)
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
