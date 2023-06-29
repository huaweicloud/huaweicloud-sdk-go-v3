package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateTemplateViewHistoriesRequest Request Object
type CreateTemplateViewHistoriesRequest struct {

	// 语言类型，缺省值为“zh-cn”。  枚举值： - zh-cn：中文 - en-us：英文
	XLanguage *CreateTemplateViewHistoriesRequestXLanguage `json:"X-Language,omitempty"`

	Body *TemplatesInfo `json:"body,omitempty"`
}

func (o CreateTemplateViewHistoriesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTemplateViewHistoriesRequest struct{}"
	}

	return strings.Join([]string{"CreateTemplateViewHistoriesRequest", string(data)}, " ")
}

type CreateTemplateViewHistoriesRequestXLanguage struct {
	value string
}

type CreateTemplateViewHistoriesRequestXLanguageEnum struct {
	ZH_CN CreateTemplateViewHistoriesRequestXLanguage
	EN_US CreateTemplateViewHistoriesRequestXLanguage
}

func GetCreateTemplateViewHistoriesRequestXLanguageEnum() CreateTemplateViewHistoriesRequestXLanguageEnum {
	return CreateTemplateViewHistoriesRequestXLanguageEnum{
		ZH_CN: CreateTemplateViewHistoriesRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: CreateTemplateViewHistoriesRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c CreateTemplateViewHistoriesRequestXLanguage) Value() string {
	return c.value
}

func (c CreateTemplateViewHistoriesRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateTemplateViewHistoriesRequestXLanguage) UnmarshalJSON(b []byte) error {
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
