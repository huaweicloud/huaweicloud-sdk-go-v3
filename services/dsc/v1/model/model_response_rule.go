package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ResponseRule 扫描规则
type ResponseRule struct {

	// 规则类别，内置规则(BUILT_IN)或自建规则(BUILT_SELF)
	Category *ResponseRuleCategory `json:"category,omitempty"`

	// 是否允许删除
	DeleteAllowed *bool `json:"delete_allowed,omitempty"`

	// 相关的规则组
	GroupNames *string `json:"group_names,omitempty"`

	// 规则ID
	Id *string `json:"id,omitempty"`

	// 逻辑运算符，\"AND\",\"OR\",\"REGEX\"
	LogicOperator *string `json:"logic_operator,omitempty"`

	// 最小匹配次数
	MinMatch *int32 `json:"min_match,omitempty"`

	// 风险等级
	RiskLevel *int32 `json:"risk_level,omitempty"`

	// 规则内容
	RuleContent *string `json:"rule_content,omitempty"`

	// 规则描述
	RuleDesc *string `json:"rule_desc,omitempty"`

	// 规则名称
	RuleName *string `json:"rule_name,omitempty"`

	// 规则类型，关键字(KEYWORD)、正则表达式(REGEX)或自然语言(NLP)
	RuleType *ResponseRuleRuleType `json:"rule_type,omitempty"`

	// 是否可选
	Selected *bool `json:"selected,omitempty"`
}

func (o ResponseRule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResponseRule struct{}"
	}

	return strings.Join([]string{"ResponseRule", string(data)}, " ")
}

type ResponseRuleCategory struct {
	value string
}

type ResponseRuleCategoryEnum struct {
	BUILT_IN   ResponseRuleCategory
	BUILT_SELF ResponseRuleCategory
}

func GetResponseRuleCategoryEnum() ResponseRuleCategoryEnum {
	return ResponseRuleCategoryEnum{
		BUILT_IN: ResponseRuleCategory{
			value: "BUILT_IN",
		},
		BUILT_SELF: ResponseRuleCategory{
			value: "BUILT_SELF",
		},
	}
}

func (c ResponseRuleCategory) Value() string {
	return c.value
}

func (c ResponseRuleCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ResponseRuleCategory) UnmarshalJSON(b []byte) error {
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

type ResponseRuleRuleType struct {
	value string
}

type ResponseRuleRuleTypeEnum struct {
	KEYWORD ResponseRuleRuleType
	REGEX   ResponseRuleRuleType
	NLP     ResponseRuleRuleType
}

func GetResponseRuleRuleTypeEnum() ResponseRuleRuleTypeEnum {
	return ResponseRuleRuleTypeEnum{
		KEYWORD: ResponseRuleRuleType{
			value: "KEYWORD",
		},
		REGEX: ResponseRuleRuleType{
			value: "REGEX",
		},
		NLP: ResponseRuleRuleType{
			value: "NLP",
		},
	}
}

func (c ResponseRuleRuleType) Value() string {
	return c.value
}

func (c ResponseRuleRuleType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ResponseRuleRuleType) UnmarshalJSON(b []byte) error {
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
