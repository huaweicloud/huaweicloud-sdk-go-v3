package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListServicePrincipalsV5Request Request Object
type ListServicePrincipalsV5Request struct {

	// 每页显示的条目数量。
	Limit *int32 `json:"limit,omitempty"`

	// 分页标记，长度为4到400个字符，只包含字母、数字、\"+\"、\"/\"、\"=\"、\"-\"和\"_\"的字符串。
	Marker *string `json:"marker,omitempty"`

	// 选择接口返回的信息的语言，可以为中文（\"zh-cn\"）或英文（\"en-us\"），默认为中文。
	XLanguage *ListServicePrincipalsV5RequestXLanguage `json:"X-Language,omitempty"`
}

func (o ListServicePrincipalsV5Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServicePrincipalsV5Request struct{}"
	}

	return strings.Join([]string{"ListServicePrincipalsV5Request", string(data)}, " ")
}

type ListServicePrincipalsV5RequestXLanguage struct {
	value string
}

type ListServicePrincipalsV5RequestXLanguageEnum struct {
	ZH_CN ListServicePrincipalsV5RequestXLanguage
	EN_US ListServicePrincipalsV5RequestXLanguage
}

func GetListServicePrincipalsV5RequestXLanguageEnum() ListServicePrincipalsV5RequestXLanguageEnum {
	return ListServicePrincipalsV5RequestXLanguageEnum{
		ZH_CN: ListServicePrincipalsV5RequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListServicePrincipalsV5RequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListServicePrincipalsV5RequestXLanguage) Value() string {
	return c.value
}

func (c ListServicePrincipalsV5RequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListServicePrincipalsV5RequestXLanguage) UnmarshalJSON(b []byte) error {
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
