package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ModifyParameterConfigTemplateRequest Request Object
type ModifyParameterConfigTemplateRequest struct {

	// 语言,默认：en-us。
	XLanguage *ModifyParameterConfigTemplateRequestXLanguage `json:"X-Language,omitempty"`

	// 参数模板ID
	ConfigId string `json:"config_id"`

	Body *ModifyParameterConfigTemplateRequestBody `json:"body,omitempty"`
}

func (o ModifyParameterConfigTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyParameterConfigTemplateRequest struct{}"
	}

	return strings.Join([]string{"ModifyParameterConfigTemplateRequest", string(data)}, " ")
}

type ModifyParameterConfigTemplateRequestXLanguage struct {
	value string
}

type ModifyParameterConfigTemplateRequestXLanguageEnum struct {
	ZH_CN ModifyParameterConfigTemplateRequestXLanguage
	EN_US ModifyParameterConfigTemplateRequestXLanguage
}

func GetModifyParameterConfigTemplateRequestXLanguageEnum() ModifyParameterConfigTemplateRequestXLanguageEnum {
	return ModifyParameterConfigTemplateRequestXLanguageEnum{
		ZH_CN: ModifyParameterConfigTemplateRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ModifyParameterConfigTemplateRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ModifyParameterConfigTemplateRequestXLanguage) Value() string {
	return c.value
}

func (c ModifyParameterConfigTemplateRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ModifyParameterConfigTemplateRequestXLanguage) UnmarshalJSON(b []byte) error {
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
