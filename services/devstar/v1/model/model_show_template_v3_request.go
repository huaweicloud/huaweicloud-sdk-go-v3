package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ShowTemplateV3Request struct {
	// 语言类型，缺省值为“zh-cn”。  枚举值： - zh-cn：中文 - en-us：英文

	XLanguage *ShowTemplateV3RequestXLanguage `json:"X-Language,omitempty"`
	// 模板ID，通过查询模板列表接口可获取相应的模板ID。

	TemplateId string `json:"template_id"`
}

func (o ShowTemplateV3Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTemplateV3Request struct{}"
	}

	return strings.Join([]string{"ShowTemplateV3Request", string(data)}, " ")
}

type ShowTemplateV3RequestXLanguage struct {
	value string
}

type ShowTemplateV3RequestXLanguageEnum struct {
	ZH_CN ShowTemplateV3RequestXLanguage
	EN_US ShowTemplateV3RequestXLanguage
}

func GetShowTemplateV3RequestXLanguageEnum() ShowTemplateV3RequestXLanguageEnum {
	return ShowTemplateV3RequestXLanguageEnum{
		ZH_CN: ShowTemplateV3RequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ShowTemplateV3RequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ShowTemplateV3RequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowTemplateV3RequestXLanguage) UnmarshalJSON(b []byte) error {
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
