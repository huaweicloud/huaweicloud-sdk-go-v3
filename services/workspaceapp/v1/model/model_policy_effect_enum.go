package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PolicyEffectEnum 状态(允许、禁用)： * `Allow` - 允许 * `Deny` - 禁用
type PolicyEffectEnum struct {
	value string
}

type PolicyEffectEnumEnum struct {
	ALLOW PolicyEffectEnum
	DENY  PolicyEffectEnum
}

func GetPolicyEffectEnumEnum() PolicyEffectEnumEnum {
	return PolicyEffectEnumEnum{
		ALLOW: PolicyEffectEnum{
			value: "Allow",
		},
		DENY: PolicyEffectEnum{
			value: "Deny",
		},
	}
}

func (c PolicyEffectEnum) Value() string {
	return c.value
}

func (c PolicyEffectEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PolicyEffectEnum) UnmarshalJSON(b []byte) error {
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
