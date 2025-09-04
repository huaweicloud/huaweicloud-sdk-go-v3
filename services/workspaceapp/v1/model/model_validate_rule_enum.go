package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ValidateRuleEnum 校验类型： * `naming` - 命名规范 * `duplicate` - 重复
type ValidateRuleEnum struct {
	value string
}

type ValidateRuleEnumEnum struct {
	NAMING    ValidateRuleEnum
	DUPLICATE ValidateRuleEnum
}

func GetValidateRuleEnumEnum() ValidateRuleEnumEnum {
	return ValidateRuleEnumEnum{
		NAMING: ValidateRuleEnum{
			value: "naming",
		},
		DUPLICATE: ValidateRuleEnum{
			value: "duplicate",
		},
	}
}

func (c ValidateRuleEnum) Value() string {
	return c.value
}

func (c ValidateRuleEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ValidateRuleEnum) UnmarshalJSON(b []byte) error {
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
