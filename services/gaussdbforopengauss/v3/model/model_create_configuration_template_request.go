package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateConfigurationTemplateRequest Request Object
type CreateConfigurationTemplateRequest struct {

	// 语言。默认值：en-us。
	XLanguage *CreateConfigurationTemplateRequestXLanguage `json:"X-Language,omitempty"`

	Body *CreateConfigurationTemplateRequestBody `json:"body,omitempty"`
}

func (o CreateConfigurationTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateConfigurationTemplateRequest struct{}"
	}

	return strings.Join([]string{"CreateConfigurationTemplateRequest", string(data)}, " ")
}

type CreateConfigurationTemplateRequestXLanguage struct {
	value string
}

type CreateConfigurationTemplateRequestXLanguageEnum struct {
	ZH_CN CreateConfigurationTemplateRequestXLanguage
	EN_US CreateConfigurationTemplateRequestXLanguage
}

func GetCreateConfigurationTemplateRequestXLanguageEnum() CreateConfigurationTemplateRequestXLanguageEnum {
	return CreateConfigurationTemplateRequestXLanguageEnum{
		ZH_CN: CreateConfigurationTemplateRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: CreateConfigurationTemplateRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c CreateConfigurationTemplateRequestXLanguage) Value() string {
	return c.value
}

func (c CreateConfigurationTemplateRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateConfigurationTemplateRequestXLanguage) UnmarshalJSON(b []byte) error {
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
