package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Template 模板名称。
type Template struct {
	value string
}

type TemplateEnum struct {
	MAGENTO   Template
	MBAAS     Template
	WORDPRESS Template
}

func GetTemplateEnum() TemplateEnum {
	return TemplateEnum{
		MAGENTO: Template{
			value: "magento",
		},
		MBAAS: Template{
			value: "mbaas",
		},
		WORDPRESS: Template{
			value: "wordpress",
		},
	}
}

func (c Template) Value() string {
	return c.value
}

func (c Template) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *Template) UnmarshalJSON(b []byte) error {
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
