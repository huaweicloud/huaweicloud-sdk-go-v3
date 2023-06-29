package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListTemplatesV2Request Request Object
type ListTemplatesV2Request struct {

	// 语言类型，缺省值为“zh-cn”。  枚举值： - zh-cn：中文 - en-us：英文
	XLanguage *ListTemplatesV2RequestXLanguage `json:"X-Language,omitempty"`

	// 请填写固定值“query”。
	ActionId string `json:"action_id"`

	Body *TemplateQueryV2 `json:"body,omitempty"`
}

func (o ListTemplatesV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTemplatesV2Request struct{}"
	}

	return strings.Join([]string{"ListTemplatesV2Request", string(data)}, " ")
}

type ListTemplatesV2RequestXLanguage struct {
	value string
}

type ListTemplatesV2RequestXLanguageEnum struct {
	ZH_CN ListTemplatesV2RequestXLanguage
	EN_US ListTemplatesV2RequestXLanguage
}

func GetListTemplatesV2RequestXLanguageEnum() ListTemplatesV2RequestXLanguageEnum {
	return ListTemplatesV2RequestXLanguageEnum{
		ZH_CN: ListTemplatesV2RequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListTemplatesV2RequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListTemplatesV2RequestXLanguage) Value() string {
	return c.value
}

func (c ListTemplatesV2RequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListTemplatesV2RequestXLanguage) UnmarshalJSON(b []byte) error {
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
