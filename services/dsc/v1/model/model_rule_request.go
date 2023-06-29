package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type RuleRequest struct {

	// 规则类别，内置规则(BUILT_IN)或自建规则(BUILT_SELF)
	Category RuleRequestCategory `json:"category"`

	// 规则ID
	Id *string `json:"id,omitempty"`

	// 逻辑运算符，\"AND\",\"OR\",\"REGEX\"
	LogicOperator string `json:"logic_operator"`

	// 最小匹配次数
	MinMatch int32 `json:"min_match"`

	// 风险等级
	RiskLevel int32 `json:"risk_level"`

	// 规则内容
	RuleContent string `json:"rule_content"`

	// 规则描述
	RuleDesc *string `json:"rule_desc,omitempty"`

	// 规则名称
	RuleName string `json:"rule_name"`

	// 规则类型，关键字(KEYWORD)、正则表达式(REGEX)或自然语言(NLP)
	RuleType RuleRequestRuleType `json:"rule_type"`
}

func (o RuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RuleRequest struct{}"
	}

	return strings.Join([]string{"RuleRequest", string(data)}, " ")
}

type RuleRequestCategory struct {
	value string
}

type RuleRequestCategoryEnum struct {
	BUILT_IN   RuleRequestCategory
	BUILT_SELF RuleRequestCategory
}

func GetRuleRequestCategoryEnum() RuleRequestCategoryEnum {
	return RuleRequestCategoryEnum{
		BUILT_IN: RuleRequestCategory{
			value: "BUILT_IN",
		},
		BUILT_SELF: RuleRequestCategory{
			value: "BUILT_SELF",
		},
	}
}

func (c RuleRequestCategory) Value() string {
	return c.value
}

func (c RuleRequestCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RuleRequestCategory) UnmarshalJSON(b []byte) error {
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

type RuleRequestRuleType struct {
	value string
}

type RuleRequestRuleTypeEnum struct {
	KEYWORD RuleRequestRuleType
	REGEX   RuleRequestRuleType
	NLP     RuleRequestRuleType
}

func GetRuleRequestRuleTypeEnum() RuleRequestRuleTypeEnum {
	return RuleRequestRuleTypeEnum{
		KEYWORD: RuleRequestRuleType{
			value: "KEYWORD",
		},
		REGEX: RuleRequestRuleType{
			value: "REGEX",
		},
		NLP: RuleRequestRuleType{
			value: "NLP",
		},
	}
}

func (c RuleRequestRuleType) Value() string {
	return c.value
}

func (c RuleRequestRuleType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RuleRequestRuleType) UnmarshalJSON(b []byte) error {
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
