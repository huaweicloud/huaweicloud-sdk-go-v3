package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListServiceSpecificCredentialsV5Request Request Object
type ListServiceSpecificCredentialsV5Request struct {

	// 选择接口返回的信息的语言，可以为中文（\"zh-cn\"）或英文（\"en-us\"），默认为中文。
	XLanguage *ListServiceSpecificCredentialsV5RequestXLanguage `json:"X-Language,omitempty"`

	// IAM用户ID。
	UserId *string `json:"user_id,omitempty"`

	// 将结果过滤为仅包含指定服务的凭证。
	ServiceName *string `json:"service_name,omitempty"`

	// 每页显示的条目数量，范围为1到200条，默认为100条。
	Limit *int32 `json:"limit,omitempty"`

	// 分页标记，长度为4到400个字符，只包含字母、数字、\"+\"、\"/\"、\"=\"、\"-\"和\"_\"的字符串。
	Marker *string `json:"marker,omitempty"`
}

func (o ListServiceSpecificCredentialsV5Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServiceSpecificCredentialsV5Request struct{}"
	}

	return strings.Join([]string{"ListServiceSpecificCredentialsV5Request", string(data)}, " ")
}

type ListServiceSpecificCredentialsV5RequestXLanguage struct {
	value string
}

type ListServiceSpecificCredentialsV5RequestXLanguageEnum struct {
	ZH_CN ListServiceSpecificCredentialsV5RequestXLanguage
	EN_US ListServiceSpecificCredentialsV5RequestXLanguage
}

func GetListServiceSpecificCredentialsV5RequestXLanguageEnum() ListServiceSpecificCredentialsV5RequestXLanguageEnum {
	return ListServiceSpecificCredentialsV5RequestXLanguageEnum{
		ZH_CN: ListServiceSpecificCredentialsV5RequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListServiceSpecificCredentialsV5RequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListServiceSpecificCredentialsV5RequestXLanguage) Value() string {
	return c.value
}

func (c ListServiceSpecificCredentialsV5RequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListServiceSpecificCredentialsV5RequestXLanguage) UnmarshalJSON(b []byte) error {
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
