package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// GetPolicyV5Request Request Object
type GetPolicyV5Request struct {

	// 身份策略ID，长度为1到64个字符，只包含字母、数字和\"-\"的字符串。
	PolicyId string `json:"policy_id"`

	// 选择接口返回的信息的语言，可以为中文（\"zh-cn\"）或英文（\"en-us\"），默认为中文。
	XLanguage *GetPolicyV5RequestXLanguage `json:"X-Language,omitempty"`
}

func (o GetPolicyV5Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetPolicyV5Request struct{}"
	}

	return strings.Join([]string{"GetPolicyV5Request", string(data)}, " ")
}

type GetPolicyV5RequestXLanguage struct {
	value string
}

type GetPolicyV5RequestXLanguageEnum struct {
	ZH_CN GetPolicyV5RequestXLanguage
	EN_US GetPolicyV5RequestXLanguage
}

func GetGetPolicyV5RequestXLanguageEnum() GetPolicyV5RequestXLanguageEnum {
	return GetPolicyV5RequestXLanguageEnum{
		ZH_CN: GetPolicyV5RequestXLanguage{
			value: "zh-cn",
		},
		EN_US: GetPolicyV5RequestXLanguage{
			value: "en-us",
		},
	}
}

func (c GetPolicyV5RequestXLanguage) Value() string {
	return c.value
}

func (c GetPolicyV5RequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GetPolicyV5RequestXLanguage) UnmarshalJSON(b []byte) error {
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
