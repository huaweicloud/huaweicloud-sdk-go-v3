package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// RuleScope 规则范围： * `PRODUCT` - 产品 * `PATH` - 路径
type RuleScope struct {
	value string
}

type RuleScopeEnum struct {
	PRODUCT RuleScope
	PATH    RuleScope
}

func GetRuleScopeEnum() RuleScopeEnum {
	return RuleScopeEnum{
		PRODUCT: RuleScope{
			value: "PRODUCT",
		},
		PATH: RuleScope{
			value: "PATH",
		},
	}
}

func (c RuleScope) Value() string {
	return c.value
}

func (c RuleScope) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RuleScope) UnmarshalJSON(b []byte) error {
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
