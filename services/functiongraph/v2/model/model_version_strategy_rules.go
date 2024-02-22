package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type VersionStrategyRules struct {

	// 参数类型。
	RuleType *VersionStrategyRulesRuleType `json:"rule_type,omitempty"`

	// 规则参数名, 只支持大小写字母，数字，下划线，中划线。
	Param *string `json:"param,omitempty"`

	// 规则匹配操作符，目前仅需支持 = 或者in。
	Op *VersionStrategyRulesOp `json:"op,omitempty"`

	// 规则值，如果op为in，则为逗号分隔的多值字符串
	Value *string `json:"value,omitempty"`
}

func (o VersionStrategyRules) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VersionStrategyRules struct{}"
	}

	return strings.Join([]string{"VersionStrategyRules", string(data)}, " ")
}

type VersionStrategyRulesRuleType struct {
	value string
}

type VersionStrategyRulesRuleTypeEnum struct {
	HEADER VersionStrategyRulesRuleType
	COOKIE VersionStrategyRulesRuleType
}

func GetVersionStrategyRulesRuleTypeEnum() VersionStrategyRulesRuleTypeEnum {
	return VersionStrategyRulesRuleTypeEnum{
		HEADER: VersionStrategyRulesRuleType{
			value: "header",
		},
		COOKIE: VersionStrategyRulesRuleType{
			value: "cookie",
		},
	}
}

func (c VersionStrategyRulesRuleType) Value() string {
	return c.value
}

func (c VersionStrategyRulesRuleType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VersionStrategyRulesRuleType) UnmarshalJSON(b []byte) error {
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

type VersionStrategyRulesOp struct {
	value string
}

type VersionStrategyRulesOpEnum struct {
	IN    VersionStrategyRulesOp
	EQUAL VersionStrategyRulesOp
}

func GetVersionStrategyRulesOpEnum() VersionStrategyRulesOpEnum {
	return VersionStrategyRulesOpEnum{
		IN: VersionStrategyRulesOp{
			value: "in",
		},
		EQUAL: VersionStrategyRulesOp{
			value: "=",
		},
	}
}

func (c VersionStrategyRulesOp) Value() string {
	return c.value
}

func (c VersionStrategyRulesOp) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VersionStrategyRulesOp) UnmarshalJSON(b []byte) error {
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
