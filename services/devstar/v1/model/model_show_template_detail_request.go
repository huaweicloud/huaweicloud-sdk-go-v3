package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ShowTemplateDetailRequest struct {

	// 语言类型，缺省值为“zh-cn”。  枚举值： - zh-cn：中文 - en-us：英文
	XLanguage *ShowTemplateDetailRequestXLanguage `json:"X-Language,omitempty" xml:"X-Language"`

	// 模板ID，通过查询模板列表接口可获取相应模板ID。
	TemplateId string `json:"template_id" xml:"template_id"`
}

func (o ShowTemplateDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTemplateDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowTemplateDetailRequest", string(data)}, " ")
}

type ShowTemplateDetailRequestXLanguage struct {
	value string
}

type ShowTemplateDetailRequestXLanguageEnum struct {
	ZH_CN ShowTemplateDetailRequestXLanguage
	EN_US ShowTemplateDetailRequestXLanguage
}

func GetShowTemplateDetailRequestXLanguageEnum() ShowTemplateDetailRequestXLanguageEnum {
	return ShowTemplateDetailRequestXLanguageEnum{
		ZH_CN: ShowTemplateDetailRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ShowTemplateDetailRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ShowTemplateDetailRequestXLanguage) Value() string {
	return c.value
}

func (c ShowTemplateDetailRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowTemplateDetailRequestXLanguage) UnmarshalJSON(b []byte) error {
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
