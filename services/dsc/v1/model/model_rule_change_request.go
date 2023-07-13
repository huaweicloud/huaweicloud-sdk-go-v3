package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type RuleChangeRequest struct {

	// 规则类别，内置规则(BUILT_IN)或自建规则(BUILT_SELF)
	Category RuleChangeRequestCategory `json:"category"`

	// 规则ID
	Id string `json:"id"`

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
	RuleType RuleChangeRequestRuleType `json:"rule_type"`
}

func (o RuleChangeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RuleChangeRequest struct{}"
	}

	return strings.Join([]string{"RuleChangeRequest", string(data)}, " ")
}

type RuleChangeRequestCategory struct {
	value string
}

type RuleChangeRequestCategoryEnum struct {
	BUILT_IN   RuleChangeRequestCategory
	BUILT_SELF RuleChangeRequestCategory
}

func GetRuleChangeRequestCategoryEnum() RuleChangeRequestCategoryEnum {
	return RuleChangeRequestCategoryEnum{
		BUILT_IN: RuleChangeRequestCategory{
			value: "BUILT_IN",
		},
		BUILT_SELF: RuleChangeRequestCategory{
			value: "BUILT_SELF",
		},
	}
}

func (c RuleChangeRequestCategory) Value() string {
	return c.value
}

func (c RuleChangeRequestCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RuleChangeRequestCategory) UnmarshalJSON(b []byte) error {
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

type RuleChangeRequestRuleType struct {
	value string
}

type RuleChangeRequestRuleTypeEnum struct {
	KEYWORD RuleChangeRequestRuleType
	REGEX   RuleChangeRequestRuleType
	NLP     RuleChangeRequestRuleType
}

func GetRuleChangeRequestRuleTypeEnum() RuleChangeRequestRuleTypeEnum {
	return RuleChangeRequestRuleTypeEnum{
		KEYWORD: RuleChangeRequestRuleType{
			value: "KEYWORD",
		},
		REGEX: RuleChangeRequestRuleType{
			value: "REGEX",
		},
		NLP: RuleChangeRequestRuleType{
			value: "NLP",
		},
	}
}

func (c RuleChangeRequestRuleType) Value() string {
	return c.value
}

func (c RuleChangeRequestRuleType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RuleChangeRequestRuleType) UnmarshalJSON(b []byte) error {
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
