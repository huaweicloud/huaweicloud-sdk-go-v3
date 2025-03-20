package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListPoliciesV5Request Request Object
type ListPoliciesV5Request struct {

	// 每页显示的条目数量。
	Limit *int32 `json:"limit,omitempty"`

	// 分页标记，长度为4到400个字符，只包含字母、数字、\"+\"、\"/\"、\"=\"、\"-\"和\"_\"的字符串。
	Marker *string `json:"marker,omitempty"`

	// 身份策略类型，可以为“自定义”（custom）或“系统预置”（system）。
	PolicyType *ListPoliciesV5RequestPolicyType `json:"policy_type,omitempty"`

	// 资源路径前缀，由若干段字符串拼接而成，每段先包含一个或多个字母、数字、\".\"、\",\"、\"+\"、\"@\"、\"=\"、\"_\"或\"-\"，并以\"/\"结尾，例如\"foo/bar/\"。
	PathPrefix *string `json:"path_prefix,omitempty"`

	// 是否仅列举存在附加实体的身份策略。
	OnlyAttached *bool `json:"only_attached,omitempty"`

	// 选择接口返回的信息的语言，可以为中文（\"zh-cn\"）或英文（\"en-us\"），默认为中文。
	XLanguage *ListPoliciesV5RequestXLanguage `json:"X-Language,omitempty"`
}

func (o ListPoliciesV5Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPoliciesV5Request struct{}"
	}

	return strings.Join([]string{"ListPoliciesV5Request", string(data)}, " ")
}

type ListPoliciesV5RequestPolicyType struct {
	value string
}

type ListPoliciesV5RequestPolicyTypeEnum struct {
	CUSTOM ListPoliciesV5RequestPolicyType
	SYSTEM ListPoliciesV5RequestPolicyType
}

func GetListPoliciesV5RequestPolicyTypeEnum() ListPoliciesV5RequestPolicyTypeEnum {
	return ListPoliciesV5RequestPolicyTypeEnum{
		CUSTOM: ListPoliciesV5RequestPolicyType{
			value: "custom",
		},
		SYSTEM: ListPoliciesV5RequestPolicyType{
			value: "system",
		},
	}
}

func (c ListPoliciesV5RequestPolicyType) Value() string {
	return c.value
}

func (c ListPoliciesV5RequestPolicyType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListPoliciesV5RequestPolicyType) UnmarshalJSON(b []byte) error {
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

type ListPoliciesV5RequestXLanguage struct {
	value string
}

type ListPoliciesV5RequestXLanguageEnum struct {
	ZH_CN ListPoliciesV5RequestXLanguage
	EN_US ListPoliciesV5RequestXLanguage
}

func GetListPoliciesV5RequestXLanguageEnum() ListPoliciesV5RequestXLanguageEnum {
	return ListPoliciesV5RequestXLanguageEnum{
		ZH_CN: ListPoliciesV5RequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListPoliciesV5RequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListPoliciesV5RequestXLanguage) Value() string {
	return c.value
}

func (c ListPoliciesV5RequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListPoliciesV5RequestXLanguage) UnmarshalJSON(b []byte) error {
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
