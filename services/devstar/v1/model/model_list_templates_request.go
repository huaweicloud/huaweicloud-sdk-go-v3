package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListTemplatesRequest struct {
	// 语言类型，缺省值为“zh-cn”。  枚举值： - zh-cn：中文 - en-us：英文

	XLanguage *ListTemplatesRequestXLanguage `json:"X-Language,omitempty"`

	Body *TemplateQuery `json:"body,omitempty"`
}

func (o ListTemplatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTemplatesRequest struct{}"
	}

	return strings.Join([]string{"ListTemplatesRequest", string(data)}, " ")
}

type ListTemplatesRequestXLanguage struct {
	value string
}

type ListTemplatesRequestXLanguageEnum struct {
	ZH_CN ListTemplatesRequestXLanguage
	EN_US ListTemplatesRequestXLanguage
}

func GetListTemplatesRequestXLanguageEnum() ListTemplatesRequestXLanguageEnum {
	return ListTemplatesRequestXLanguageEnum{
		ZH_CN: ListTemplatesRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListTemplatesRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListTemplatesRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListTemplatesRequestXLanguage) UnmarshalJSON(b []byte) error {
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
