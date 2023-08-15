package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type Rule struct {

	// 回收类型，date_rule、tag_rule
	Template RuleTemplate `json:"template"`

	// template是date_rule时，设置params为{\"days\": \"xxx\"} template是tag_rule时，设置params为{\"num\": \"xxx\"}
	Params *interface{} `json:"params"`

	// 例外镜像
	TagSelectors []TagSelector `json:"tag_selectors"`
}

func (o Rule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Rule struct{}"
	}

	return strings.Join([]string{"Rule", string(data)}, " ")
}

type RuleTemplate struct {
	value string
}

type RuleTemplateEnum struct {
	DATE_RULE RuleTemplate
	TAG_RULE  RuleTemplate
}

func GetRuleTemplateEnum() RuleTemplateEnum {
	return RuleTemplateEnum{
		DATE_RULE: RuleTemplate{
			value: "date_rule",
		},
		TAG_RULE: RuleTemplate{
			value: "tag_rule",
		},
	}
}

func (c RuleTemplate) Value() string {
	return c.value
}

func (c RuleTemplate) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RuleTemplate) UnmarshalJSON(b []byte) error {
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
