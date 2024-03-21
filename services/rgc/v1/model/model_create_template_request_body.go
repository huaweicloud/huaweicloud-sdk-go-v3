package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateTemplateRequestBody 创建模板请求参数。
type CreateTemplateRequestBody struct {

	// 模板名称。
	TemplateName string `json:"template_name"`

	// 模板类型，包括预置和自定义。
	TemplateType CreateTemplateRequestBodyTemplateType `json:"template_type"`

	// 模板描述。
	TemplateDescription *string `json:"template_description,omitempty"`

	// 模板内容。
	TemplateBody *string `json:"template_body,omitempty"`

	// 模板文件名。
	TemplateFileName string `json:"template_file_name"`
}

func (o CreateTemplateRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTemplateRequestBody struct{}"
	}

	return strings.Join([]string{"CreateTemplateRequestBody", string(data)}, " ")
}

type CreateTemplateRequestBodyTemplateType struct {
	value string
}

type CreateTemplateRequestBodyTemplateTypeEnum struct {
	PREDEFINED CreateTemplateRequestBodyTemplateType
	CUSTOMIZED CreateTemplateRequestBodyTemplateType
}

func GetCreateTemplateRequestBodyTemplateTypeEnum() CreateTemplateRequestBodyTemplateTypeEnum {
	return CreateTemplateRequestBodyTemplateTypeEnum{
		PREDEFINED: CreateTemplateRequestBodyTemplateType{
			value: "predefined",
		},
		CUSTOMIZED: CreateTemplateRequestBodyTemplateType{
			value: "customized",
		},
	}
}

func (c CreateTemplateRequestBodyTemplateType) Value() string {
	return c.value
}

func (c CreateTemplateRequestBodyTemplateType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateTemplateRequestBodyTemplateType) UnmarshalJSON(b []byte) error {
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
