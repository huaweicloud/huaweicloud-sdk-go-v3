package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 结构化类型。
type ShowStructTemplateRule struct {
	// 测试

	Param *string `json:"param,omitempty"`
	// 结构化类型

	Type *ShowStructTemplateRuleType `json:"type,omitempty"`
}

func (o ShowStructTemplateRule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowStructTemplateRule struct{}"
	}

	return strings.Join([]string{"ShowStructTemplateRule", string(data)}, " ")
}

type ShowStructTemplateRuleType struct {
	value string
}

type ShowStructTemplateRuleTypeEnum struct {
	JSON         ShowStructTemplateRuleType
	SPLIT        ShowStructTemplateRuleType
	NGINX        ShowStructTemplateRuleType
	BUILT_IN     ShowStructTemplateRuleType
	CUSTOM_REGEX ShowStructTemplateRuleType
}

func GetShowStructTemplateRuleTypeEnum() ShowStructTemplateRuleTypeEnum {
	return ShowStructTemplateRuleTypeEnum{
		JSON: ShowStructTemplateRuleType{
			value: "json",
		},
		SPLIT: ShowStructTemplateRuleType{
			value: "split",
		},
		NGINX: ShowStructTemplateRuleType{
			value: "nginx",
		},
		BUILT_IN: ShowStructTemplateRuleType{
			value: "built_in",
		},
		CUSTOM_REGEX: ShowStructTemplateRuleType{
			value: "custom_regex",
		},
	}
}

func (c ShowStructTemplateRuleType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowStructTemplateRuleType) UnmarshalJSON(b []byte) error {
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
