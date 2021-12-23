package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 创建转发规则的请求体
type CreateL7ruleReqInPolicy struct {
	// 转发规则的管理状态；该字段为预留字段，暂未启用。默认为true。

	AdminStateUp *bool `json:"admin_state_up,omitempty"`
	// 转发规则的匹配内容

	Type CreateL7ruleReqInPolicyType `json:"type"`
	// 转发规则的匹配方式。type为HOST_NAME时可以为EQUAL_TO。type为PATH时可以为REGEX， STARTS_WITH，EQUAL_TO。

	CompareType string `json:"compare_type"`
	// 匹配内容的键值。目前匹配内容为HOST_NAME和PATH时，该字段不生效。该字段能更新但不会生效。

	Key *string `json:"key,omitempty"`
	// 匹配内容的值。其值不能包含空格。使用说明：当type为HOST_NAME时，取值范围：String(100)，字符串只能包含英文字母、数字、“-”或“.”，且必须以字母或数字开头。当type为PATH时，取值范围：String(128)。当转发规则的compare_type为STARTS_WITH，EQUAL_TO时，字符串只能包含英文字母、数字、_~';@^-%#&$.*+?,=!:| /()[]{}，且必须以\"/\"开头。

	Value string `json:"value"`
	// 是否反向匹配；取值范围：true/false。默认值：false；该字段为预留字段，暂未启用。

	Invert *bool `json:"invert,omitempty"`
}

func (o CreateL7ruleReqInPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateL7ruleReqInPolicy struct{}"
	}

	return strings.Join([]string{"CreateL7ruleReqInPolicy", string(data)}, " ")
}

type CreateL7ruleReqInPolicyType struct {
	value string
}

type CreateL7ruleReqInPolicyTypeEnum struct {
	HOST_NAME CreateL7ruleReqInPolicyType
	PATH      CreateL7ruleReqInPolicyType
}

func GetCreateL7ruleReqInPolicyTypeEnum() CreateL7ruleReqInPolicyTypeEnum {
	return CreateL7ruleReqInPolicyTypeEnum{
		HOST_NAME: CreateL7ruleReqInPolicyType{
			value: "HOST_NAME",
		},
		PATH: CreateL7ruleReqInPolicyType{
			value: "PATH",
		},
	}
}

func (c CreateL7ruleReqInPolicyType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateL7ruleReqInPolicyType) UnmarshalJSON(b []byte) error {
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
