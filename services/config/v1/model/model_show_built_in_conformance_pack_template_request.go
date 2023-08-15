package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowBuiltInConformancePackTemplateRequest Request Object
type ShowBuiltInConformancePackTemplateRequest struct {

	// 合规规则包模板ID。
	TemplateId string `json:"template_id"`

	// 选择接口返回的信息的语言，默认为\"zh-cn\"中文
	XLanguage *ShowBuiltInConformancePackTemplateRequestXLanguage `json:"X-Language,omitempty"`
}

func (o ShowBuiltInConformancePackTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBuiltInConformancePackTemplateRequest struct{}"
	}

	return strings.Join([]string{"ShowBuiltInConformancePackTemplateRequest", string(data)}, " ")
}

type ShowBuiltInConformancePackTemplateRequestXLanguage struct {
	value string
}

type ShowBuiltInConformancePackTemplateRequestXLanguageEnum struct {
	ZH_CN ShowBuiltInConformancePackTemplateRequestXLanguage
	EN_US ShowBuiltInConformancePackTemplateRequestXLanguage
}

func GetShowBuiltInConformancePackTemplateRequestXLanguageEnum() ShowBuiltInConformancePackTemplateRequestXLanguageEnum {
	return ShowBuiltInConformancePackTemplateRequestXLanguageEnum{
		ZH_CN: ShowBuiltInConformancePackTemplateRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ShowBuiltInConformancePackTemplateRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ShowBuiltInConformancePackTemplateRequestXLanguage) Value() string {
	return c.value
}

func (c ShowBuiltInConformancePackTemplateRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowBuiltInConformancePackTemplateRequestXLanguage) UnmarshalJSON(b []byte) error {
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
